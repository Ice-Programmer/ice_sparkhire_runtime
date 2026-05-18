package search

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/model/db"
	elastic "ice_sparkhire_runtime/model/elasticsearch"
	"ice_sparkhire_runtime/utils"
)

type ESSearchStrategy struct{}

func NewESSearchStrategy() *ESSearchStrategy {
	return &ESSearchStrategy{}
}

func (s *ESSearchStrategy) Search(ctx context.Context, searchParam *SearchParam) (*SearchResult, error) {
	recruitmentDocList, total, err := elastic.SearchRecruitmentDoc(ctx, searchParam.PageSize, searchParam.PageNum, searchParam.Condition)
	if err != nil {
		return nil, err
	}

	ids := utils.MapStructList(recruitmentDocList, func(recruitmentDoc *elastic.RecruitmentDoc) int64 {
		return recruitmentDoc.ID
	})

	if len(ids) == 0 {
		return &SearchResult{
			Total:      0,
			ResultList: []*db.Recruitment{},
		}, nil
	}

	// search in db
	recruitmentDBList, err := db.ListRecruitmentByIds(ctx, db.DB, ids)
	if err != nil {
		return nil, fmt.Errorf("mysql mget failed: %w", err)
	}

	// 按 ES 返回的 ID 顺序进行内存重排
	rowMap := make(map[int64]*db.Recruitment, len(recruitmentDBList))
	for _, row := range recruitmentDBList {
		rowMap[row.ID] = row
	}

	orderedResult := make([]*db.Recruitment, 0, len(ids))
	for _, id := range ids {
		if row, exists := rowMap[id]; exists {
			orderedResult = append(orderedResult, row)
		}
	}

	return &SearchResult{
		Total:      total,
		ResultList: orderedResult,
	}, nil
}
