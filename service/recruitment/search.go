package recruitment

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/service/milvus"
)

func SearchFromEmbedding(ctx context.Context, queryText string, topK int) ([]*RecruitmentMilvusItem, error) {
	milvusManager := milvus.NewMilvusManager(ctx, RecruitmentCollectionSchema, RecruitmentCollectionName)
	returnCols := []string{RecruitmentIdCol, milvus.ScoreCol}

	// 2. search
	normalSearchResults, err := milvusManager.Search(queryText, "", topK)
	if err != nil {
		klog.CtxErrorf(ctx, "milvus Normal Search err: %v", err)
		return nil, err
	}

	recruitmentItemList, err := milvus.ParseSearchResult[RecruitmentMilvusItem](ctx, normalSearchResults, RecruitmentCollectionSchema, returnCols...)
	if err != nil {
		klog.CtxErrorf(ctx, "milvus Normal Parse err: %v", err)
		return nil, err
	}

	return recruitmentItemList, nil
}
