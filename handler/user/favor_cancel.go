package user

import (
	"context"
	"fmt"
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

	// fetch favor record
	favor, err := db.FindUserFavorByTargetIdAndUserId(ctx, db.DB, userId, req.TargetId, req.TargetType)
	if err != nil {
		return nil, fmt.Errorf("find user favorite err: %v", err)
	}

	// delete favor record
	if err := db.DeleteUserFavor(ctx, db.DB, favor.Id); err != nil {
		return nil, fmt.Errorf("delete user favorite err: %v", err)
	}

	return &sparkruntime.UserCancelFavorResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
