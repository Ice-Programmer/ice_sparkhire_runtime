package milvus

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/service/model"
	"ice_sparkhire_runtime/utils"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

const DefaultDim = 2048

type MilvusManager struct {
	Ctx            context.Context
	Schema         *MilvusSchema
	CollectionName string
}

func NewMilvusManager(ctx context.Context, schema *MilvusSchema, collectionName string) *MilvusManager {
	return &MilvusManager{
		Ctx:            ctx,
		Schema:         schema,
		CollectionName: collectionName,
	}
}

func (h *MilvusManager) CreateCollection() error {
	exist, err := milvusClient.HasCollection(h.Ctx, h.CollectionName)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Check collection %s Exist fail, err: %v", h.CollectionName, err)
	}
	if exist {
		if err := milvusClient.DropCollection(h.Ctx, h.CollectionName); err != nil {
			klog.CtxErrorf(h.Ctx, "[Milvus Service] Drop collection %s fail, err: %v", h.CollectionName, err)
			return err
		}
	}

	schema := h.Schema.GenCollectionSchema(h.Ctx, h.CollectionName)

	if err = milvusClient.CreateCollection(h.Ctx, schema, entity.DefaultShardNumber); err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Create collection %s fail, err: %v", h.CollectionName, err)
		return err
	}

	klog.CtxInfof(h.Ctx, "[Milvus Service] Create collection %s success", h.CollectionName)
	return nil
}

func (h *MilvusManager) CreateHNSWIndex(M, ef int) error {
	index, err := entity.NewIndexHNSW(entity.L2, M, ef)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Create HNSW Index fail, collection: %s, err: %v", h.CollectionName, err)
	}

	vectorColName := h.Schema.GetFirstVectorColName()

	if err = milvusClient.CreateIndex(h.Ctx, h.CollectionName, vectorColName, index, false); err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Create HNSW Index fail, collection name: %s, err: %v", h.CollectionName, err)
		return err
	}

	klog.CtxInfof(h.Ctx, "[Milvus Service] Create HNSW Index successfully, collection: %s, col: %s", h.CollectionName, vectorColName)
	return nil
}

func (h *MilvusManager) InsertData(column ...entity.Column) error {
	idCol, err := milvusClient.Insert(h.Ctx, h.CollectionName, "", column...)
	if err != nil {
		klog.CtxInfof(h.Ctx, "[Milvus Service] Insert fail, err: %v", h.CollectionName)
		return err
	}
	klog.CtxInfof(h.Ctx, "[Milvus Service] Insert data successfully, collection: %s, col num: %s", h.CollectionName, idCol.Len())
	return nil
}

func (h *MilvusManager) UpsertData(column ...entity.Column) error {
	idCol, err := milvusClient.Upsert(h.Ctx, h.CollectionName, "", column...)
	if err != nil {
		klog.CtxInfof(h.Ctx, "[Milvus Service] Insert fail, err: %v", h.CollectionName)
		return err
	}
	klog.CtxInfof(h.Ctx, "[Milvus Service] Insert data successfully, collection: %s, col num: %s", h.CollectionName, idCol.Len())
	return nil
}

func (h *MilvusManager) RemoveData(expr string) error {
	if err := milvusClient.Delete(h.Ctx, h.CollectionName, "", expr); err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Delete fail, collectionName: %s, expr: %s, err: %v", h.CollectionName, expr, err)
		return err
	}

	return nil
}

func (h *MilvusManager) FlushAndLoad() error {
	if err := milvusClient.Flush(h.Ctx, h.CollectionName, false); err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Flush %s fail, err: %v", h.CollectionName, err)
		return err
	}

	if err := milvusClient.LoadCollection(h.Ctx, h.CollectionName, false); err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Load %s fail, err: %v", h.CollectionName, err)
		return err
	}
	return nil
}

func (h *MilvusManager) Search(query string, filterExpr string, topK int, returnCols ...string) ([]client.SearchResult, error) {
	embeddedQueryText, err := model.EmbeddingText(h.Ctx, query)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Search fail, err: %v", err.Error())
		return nil, err
	}

	if len(embeddedQueryText) != DefaultDim {
		return nil, fmt.Errorf("query vector dim mismatch, expect %d", DefaultDim)
	}

	vectors := []entity.Vector{
		entity.FloatVector(embeddedQueryText),
	}

	if len(returnCols) == 0 {
		returnCols = h.Schema.GetColNameList()
	}

	searchParam, err := entity.NewIndexHNSWSearchParam(300)
	if err != nil {
		return nil, err
	}

	searchResults, err := milvusClient.Search(
		h.Ctx,
		h.CollectionName,
		[]string{},
		filterExpr,
		returnCols,
		vectors,
		h.Schema.GetFirstVectorColName(),
		entity.L2,
		topK,
		searchParam,
	)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Search fail, err: %v", err.Error())
		return nil, err
	}

	return searchResults, nil
}

func (h *MilvusManager) SearchByPage(query string, filterExpr string, pageNum, pageSize int32, returnCols ...string) ([]client.SearchResult, error) {
	//embeddedQueryText, err := .GenerateTextEmbedding(h.Ctx, query)
	embeddedQueryText, err := model.EmbeddingText(h.Ctx, query)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Search fail, err: %v", err.Error())
		return nil, err
	}

	if len(embeddedQueryText) != DefaultDim {
		return nil, fmt.Errorf("query vector dim mismatch, expect %d", DefaultDim)
	}

	vectors := []entity.Vector{
		entity.FloatVector(embeddedQueryText),
	}

	if len(returnCols) == 0 {
		returnCols = h.Schema.GetColNameList()
	}

	searchParam, err := entity.NewIndexHNSWSearchParam(300)
	if err != nil {
		return nil, err
	}

	pageSize, pageNum = utils.SetPageDefault(pageSize, pageNum)

	searchResults, err := milvusClient.Search(
		h.Ctx,
		h.CollectionName,
		[]string{},
		filterExpr,
		returnCols,
		vectors,
		h.Schema.GetFirstVectorColName(),
		entity.L2,
		int(pageSize*pageNum),
		searchParam,
		client.WithLimit(int64(pageSize)),
		client.WithOffset(int64(utils.GetPageOffset(pageNum, pageSize))),
	)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Search fail, err: %v", err.Error())
		return nil, err
	}

	return searchResults, nil
}

func (h *MilvusManager) Query(filterExpr string, returnColumns []string) (client.ResultSet, error) {
	if len(returnColumns) == 0 {
		returnColumns = h.Schema.GetColNameList()
	}
	queryResults, err := milvusClient.Query(
		h.Ctx,
		h.CollectionName,
		nil,
		filterExpr,
		returnColumns,
		client.WithLimit(1000),
	)
	if err != nil {
		klog.CtxErrorf(h.Ctx, "[Milvus Service] Query fail, err: %v", err.Error())
		return nil, err
	}

	return queryResults, nil
}

func GetDefaultTopK(topK int64) int {
	if topK <= 0 || topK > consts.MaxTopK {
		return consts.DefaultTopK
	}

	return int(topK)
}
