package search

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	elastic "ice_sparkhire_runtime/model/elasticsearch"
)

type SearchParam struct {
	PageSize  int32
	PageNum   int32
	Condition *sparkruntime.RecruitmentCondition
}

type SearchResult struct {
	Total      int64
	ResultList []*db.Recruitment
}

type SearchStrategy interface {
	Search(ctx context.Context, searchParam *SearchParam) (*SearchResult, error)
}

type SearchFactory struct{}

func NewSearchFactory() *SearchFactory {
	return &SearchFactory{}
}

// GetSearchEngine 根据 ES 客户端状态动态获取搜索策略
func (f *SearchFactory) GetSearchEngine() SearchStrategy {
	if elastic.ValidClient() {
		return NewESSearchStrategy()
	}
	return NewDBSearchStrategy()
}
