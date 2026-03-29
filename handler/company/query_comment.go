package company

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/company"
	"ice_sparkhire_runtime/utils"
)

func QueryCompanyCommentPage(ctx context.Context, req *sparkruntime.QueryCompanyCommentPageRequest) (*sparkruntime.QueryCompanyCommentPageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	// find root comment
	commentList, total, err := db.QueryRootCompanyCommentPage(ctx, db.DB, pageSize, pageNum, req.GetCondition())
	if err != nil {
		return nil, err
	}

	commentInfoList, err := company.BuildCommentInfoList(ctx, commentList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.QueryCompanyCommentPageResponse{
		CommentInfoList: commentInfoList,
		Total:           total,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}
