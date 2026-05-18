package elasticsearch

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v9/esutil"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
	"strconv"
	"time"
)

const RecruitmentIndexName = "recruitment"

type RecruitmentDoc struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	CompanyID     int64     `json:"company_id"`
	CareerID      int64     `json:"career_id"`
	Description   string    `json:"description"`
	Requirement   string    `json:"requirement"`
	JobType       int8      `json:"job_type"`
	EducationType int8      `json:"education_type"`
	SalaryLower   int32     `json:"salary_lower"`
	SalaryUpper   int32     `json:"salary_upper"`
	Address       string    `json:"address"`
	Location      *GeoPoint `json:"location,omitempty"`
	Status        int8      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func FromRecruitmentDB(recruitment *db.Recruitment) *RecruitmentDoc {
	return &RecruitmentDoc{
		ID:            recruitment.ID,
		Name:          recruitment.Name,
		CompanyID:     recruitment.CompanyId,
		CareerID:      recruitment.CareerId,
		Description:   recruitment.Description,
		Requirement:   recruitment.Requirement,
		JobType:       recruitment.JobType,
		EducationType: recruitment.EducationType,
		SalaryLower:   recruitment.SalaryLower,
		SalaryUpper:   recruitment.SalaryUpper,
		Address:       recruitment.Address,
		Location: &GeoPoint{
			Lat: recruitment.Latitude,
			Lon: recruitment.Longitude,
		},
		Status:    recruitment.Status,
		CreatedAt: recruitment.CreatedAt,
	}
}

func CreateRecruitmentIndexDoc(ctx context.Context, doc RecruitmentDoc) error {
	docID := strconv.FormatInt(doc.ID, 10)
	// 使用 Typed API 进行 Index 操
	res, err := elasticClient.Index(RecruitmentIndexName).
		Id(docID).
		Request(doc).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("error indexing document ID=%d: %w", doc.ID, err)
	}

	klog.CtxInfof(ctx, "Document indexed successfully. Result: %s, Version: %d\n", res.Result, res.Version_)
	return nil
}

func BulkCreateRecruitmentDocs(ctx context.Context, docs []*RecruitmentDoc) error {
	if len(docs) == 0 {
		return nil
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        elasticClient,        // 基础客户端
		Index:         RecruitmentIndexName, // 目标索引名
		NumWorkers:    4,                    // 并发协程数
		FlushBytes:    5e+6,                 // 达到约 5MB 时自动刷新提交
		FlushInterval: 5 * time.Second,      // 或每隔 5 秒强制刷新提交
	})
	if err != nil {
		return fmt.Errorf("error creating bulk indexer: %w", err)
	}

	for _, doc := range docs {
		docID := strconv.FormatInt(doc.ID, 10)

		data, err := sonic.Marshal(doc)
		if err != nil {
			return fmt.Errorf("error marshaling doc ID=%d: %w", doc.ID, err)
		}

		err = bi.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: docID,
				Body:       bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					klog.CtxInfof(ctx, "Document ID=%s indexed successfully. Version: %d", item.DocumentID, res.Version)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						klog.CtxErrorf(ctx, "ERROR indexing document ID=%s: %s", item.DocumentID, err)
					} else {
						klog.CtxErrorf(ctx, "ERROR indexing document ID=%s: [%s] %s", item.DocumentID, res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			return fmt.Errorf("bulk indexer add error: %w", err)
		}
	}

	if err := bi.Close(ctx); err != nil {
		return fmt.Errorf("bulk indexer close error: %w", err)
	}

	stats := bi.Stats()
	klog.CtxInfof(ctx, "Bulk index total finished. Added: %d, Flushed: %d, Failed: %d", stats.NumAdded, stats.NumFlushed, stats.NumFailed)

	if stats.NumFailed > 0 {
		return fmt.Errorf("bulk index completed with %d failures", stats.NumFailed)
	}

	return nil
}

func DeleteRecruitmentIndexDoc(ctx context.Context, recruitmentId int64) error {
	docID := strconv.FormatInt(recruitmentId, 10)
	if _, err := elasticClient.Delete(RecruitmentIndexName, docID).Do(ctx); err != nil {
		return fmt.Errorf("error deleting recruitment ID=%d: %w", recruitmentId, err)
	}

	return nil
}

func SyncAllRecruitmentToES(ctx context.Context) error {
	const batchSize = 500
	var lastID int64 = 0
	totalSynced := 0

	for {
		recruitments, err := db.FindRecruitmentBatchByCursor(ctx, db.DB, lastID, batchSize)
		if err != nil {
			return err
		}
		if len(recruitments) == 0 {
			break
		}

		var docs []*RecruitmentDoc
		for _, item := range recruitments {
			docs = append(docs, FromRecruitmentDB(item))
			lastID = item.ID
		}

		if err := BulkCreateRecruitmentDocs(ctx, docs); err != nil {
			klog.CtxErrorf(ctx, "[ES_SYNC] Batch failed at lastID=%d: %v", lastID, err)
			return err
		}

		totalSynced += len(recruitments)
	}
	return nil
}

func SearchRecruitmentDoc(ctx context.Context, pageSize, pageNum int32, condition *sparkruntime.RecruitmentCondition) ([]*RecruitmentDoc, int64, error) {
	mustQueries, filterQueries, err := buildRecruitmentSearchParam(condition)
	if err != nil {
		return nil, 0, err
	}

	from := int((pageNum - 1) * pageSize)

	res, err := elasticClient.Search().
		Index(RecruitmentIndexName).
		Request(&search.Request{
			From: utils.IntPtr(from),
			Size: utils.IntPtr(int(pageSize)),
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must:   mustQueries,
					Filter: filterQueries,
				},
			},
			Sort: []types.SortCombinations{
				types.SortOptions{
					SortOptions: map[string]types.FieldSort{
						"created_at": {Order: &sortorder.Desc},
					},
				},
			},
		}).
		Do(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("es search failed: %w", err)
	}

	recruitmentDocList := make([]*RecruitmentDoc, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var doc RecruitmentDoc
		if err := sonic.Unmarshal(hit.Source_, &doc); err != nil {
			return nil, 0, err
		}
		recruitmentDocList = append(recruitmentDocList, &doc)
	}

	return recruitmentDocList, res.Hits.Total.Value, nil
}

func buildRecruitmentSearchParam(condition *sparkruntime.RecruitmentCondition) (mustQueries []types.Query, filterQueries []types.Query, err error) {
	if condition == nil {
		return mustQueries, filterQueries, nil
	}

	// 关键词检索
	if condition.IsSetSearchText() && len(condition.GetSearchText()) > 0 {
		mustQueries = append(mustQueries, types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query:  condition.GetSearchText(),
				Fields: []string{"name^2", "description", "requirement"},
			},
		})
	}

	// 过滤公司 id
	if condition.IsSetCompanyId() {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"company_id": {Value: condition.GetCompanyId()},
			},
		})
	}

	// 过滤 job type
	if condition.IsSetCareerId() {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"career_id": {Value: condition.GetCareerId()},
			},
		})
	}

	// 过滤 salary upper
	if condition.IsSetSalaryLower() {
		filterQueries = append(filterQueries, types.Query{
			Range: map[string]types.RangeQuery{
				"salary_upper": types.LongNumberRangeQuery{
					Lte: utils.Int64Ptr(int64(condition.GetSalaryUpper())),
				},
			},
		})
	}

	// 过滤 salary lower
	if condition.IsSetSalaryLower() {
		filterQueries = append(filterQueries, types.Query{
			Range: map[string]types.RangeQuery{
				"salary_lower": types.LongNumberRangeQuery{
					Gte: utils.Int64Ptr(int64(condition.GetSalaryLower())),
				},
			},
		})
	}

	// 过滤 education status
	if condition.IsSetEducationStatus() {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"education_type": {Value: condition.GetEducationStatus()},
			},
		})
	}

	return mustQueries, filterQueries, nil
}
