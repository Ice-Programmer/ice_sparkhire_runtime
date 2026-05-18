package search

import (
	"context"
	"ice_sparkhire_runtime/model/db"
)

type DBSearchStrategy struct {
}

func NewDBSearchStrategy() *DBSearchStrategy {
	return &DBSearchStrategy{}
}

func (s *DBSearchStrategy) Search(ctx context.Context, params *SearchParam) (*SearchResult, error) {
	recruitmentList, total, err := db.QueryRecruitmentPage(ctx, db.DB, params.PageSize, params.PageNum, params.Condition)
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		Total:      total,
		ResultList: recruitmentList,
	}, nil
}
