package user

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func UserFavor(ctx context.Context, req *sparkruntime.UserFavorRequest) (*sparkruntime.UserFavorResponse, error) {
	if req.GetTargetId() <= 0 {
		return nil, fmt.Errorf("target id is invalid")
	}

	if !utils.IsValidEnum(req.GetTargetType()) {
		return nil, fmt.Errorf("target type is invalid")
	}

	// fetch current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current user id error: %w", err)
	}

	err = db.DB.Transaction(func(tx *gorm.DB) (err error) {
		// 1. 判断是否已经点赞
		hasFavor := db.HasFavor(ctx, tx, userId, req.TargetId, req.TargetType)
		if hasFavor {
			return fmt.Errorf("duplicate favor operation")
		}

		// 2. create user favor
		if err := db.CreateUserFavor(ctx, tx, &db.UserFavorite{
			UserId:     userId,
			TargetId:   req.TargetId,
			TargetType: int8(req.TargetType),
		}); err != nil {
			return err
		}

		// 3. 计数 +1
		switch req.TargetType {
		case sparkruntime.TargetType_Recruitment:
			// todo
		case sparkruntime.TargetType_Company:
			err = db.UpdateCompanyFavorCnt(ctx, tx, req.TargetId, +1)
		}
		if err != nil {
			return fmt.Errorf("find %s error", req.GetTargetType().String())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &sparkruntime.UserFavorResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
