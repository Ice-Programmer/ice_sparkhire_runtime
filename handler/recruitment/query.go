package recruitment

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/recruitment"
	"ice_sparkhire_runtime/service/recruitment/search"
	"ice_sparkhire_runtime/utils"
	"time"
)

func QueryRecruitmentPage(ctx context.Context, req *sparkruntime.QueryRecruitmentPageRequest) (*sparkruntime.QueryRecruitmentPageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

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

	if req.IsSetCondition() && len(req.GetCondition().GetSearchText()) > 0 {
		// 记录搜索记录
		err := db.CreateUserSearchHistory(ctx, db.DB, &db.UserSearchHistory{
			Id:            utils.GetId(),
			UserId:        userId,
			SearchContent: req.GetCondition().GetSearchText(),
			Type:          sparkruntime.UserHistoryType_Recruitment,
			CreatedAt:     time.Now(),
		})
		if err != nil {
			klog.CtxErrorf(ctx, "record user history error: %v", err)
		}
	}

	return &sparkruntime.QueryRecruitmentPageResponse{
		RecruitmentList: recruitmentPageInfos,
		Total:           searchResult.Total,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}
