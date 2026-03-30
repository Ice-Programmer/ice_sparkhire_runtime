package company

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func DeleteCompanyComment(ctx context.Context, req *sparkruntime.DeleteCompanyCommentRequest) (*sparkruntime.DeleteCompanyCommentResponse, error) {
	if req.GetId() <= 0 {
		return nil, fmt.Errorf("comment id is invalid")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	comment, err := db.FindCompanyCommentById(ctx, db.DB, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("find comment error: %v", err)
	}

	if comment.UserId != userId {
		return nil, fmt.Errorf("only can delete own comment")
	}

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 判断是否为 reply comment
		if comment.RootId != 0 {
			if err := db.UpdateCompanyCommentReplyCount(ctx, tx, comment.RootId, -1); err != nil {
				return err
			}
		}

		// delete comment
		if err := db.DeleteCompanyCommentById(ctx, tx, req.Id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.DeleteCompanyCommentResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
