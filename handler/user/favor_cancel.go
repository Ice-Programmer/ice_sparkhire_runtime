package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func UserCancelFavor(ctx context.Context, req *sparkruntime.UserCancelFavorRequest) (*sparkruntime.UserCancelFavorResponse, error) {
	if req.GetTargetId() <= 0 {
		return nil, fmt.Errorf("invalid target id")
	}

	if !utils.IsValidEnum(req.GetTargetType()) {
		return nil, fmt.Errorf("invalid target type")
	}

	// fetch current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current user info err: %v", err)
	}

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		// delete favor record
		if err := db.DeleteUserFavor(ctx, tx, userId, req.TargetId, req.TargetType); err != nil {
			return fmt.Errorf("delete user favorite err: %v", err)
		}

		// update favorite count
		switch req.TargetType {
		case sparkruntime.TargetType_Recruitment:
			// todo
		case sparkruntime.TargetType_Company:
			err = db.UpdateCompanyFavorCnt(ctx, tx, req.TargetId, -1)
		}
		if err != nil {
			return fmt.Errorf("find %s error", req.GetTargetType().String())
		}

		return nil
	}); err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] delete userFavorite error: %v", err)
		return nil, err
	}

	return &sparkruntime.UserCancelFavorResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
