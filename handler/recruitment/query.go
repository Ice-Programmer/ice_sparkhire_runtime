package recruitment

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/recruitment"
	"ice_sparkhire_runtime/service/recruitment/search"
	"ice_sparkhire_runtime/utils"
)

func QueryRecruitmentPage(ctx context.Context, req *sparkruntime.QueryRecruitmentPageRequest) (*sparkruntime.QueryRecruitmentPageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	searchEngine := search.NewSearchFactory().
		GetSearchEngine()

	searchResult, err := searchEngine.Search(ctx, &search.SearchParam{
		PageSize:  pageSize,
		PageNum:   pageNum,
		Condition: req.GetCondition(),
	})
	if err != nil {
		return nil, err
	}

	recruitmentPageInfos, err := recruitment.BuildEvaluatePageInfo(ctx, searchResult.ResultList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.QueryRecruitmentPageResponse{
		RecruitmentList: recruitmentPageInfos,
		Total:           searchResult.Total,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}
