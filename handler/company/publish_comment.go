package company

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/company"
	"ice_sparkhire_runtime/utils"
)

func PublishCompanyComment(ctx context.Context, req *sparkruntime.PublishCompanyCommentRequest) (*sparkruntime.PublishCompanyCommentResponse, error) {
	if err := validateCompanyComment(ctx, req); err != nil {
		return nil, err
	}

	// fetch current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	comment := &db.CompanyComment{
		Id:        utils.GetId(),
		CompanyId: req.CompanyId,
		UserId:    userId,
		Content:   req.Content,
	}

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		// 判断是否为 reply comment
		if req.IsSetParentId() {
			// 为回复评论
			comment, err = company.BuildReplyComment(ctx, req.GetParentId(), comment)
			if err != nil {
				return err
			}

			// update reply num
			if err := db.IncrCompanyCommentReplyCount(ctx, tx, comment.RootId); err != nil {
				return err
			}
		}

		// create comment
		if err := db.CreateCompanyComment(ctx, tx, comment); err != nil {
			return fmt.Errorf("create comment error: %v", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.PublishCompanyCommentResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateCompanyComment(ctx context.Context, req *sparkruntime.PublishCompanyCommentRequest) error {
	if req.GetCompanyId() <= 0 {
		return fmt.Errorf("invalid company id")
	}

	if len(req.GetContent()) == 0 {
		return fmt.Errorf("comment content can not be empty")
	}

	if len(req.GetContent()) > 500 {
		return fmt.Errorf("comment content can not be longer than 500")
	}

	// company exist
	if _, err := db.FindCompanyById(ctx, db.DB, req.CompanyId); err != nil {
		return fmt.Errorf("find company error: %v", err)
	}

	return nil
}
