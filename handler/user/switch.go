package user

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func SwitchUserRole(ctx context.Context, req *sparkruntime.SwitchUserRoleRequest) (*sparkruntime.SwitchUserRoleResponse, error) {
	if req.GetUserRole() != sparkruntime.UserRole_HR && req.GetUserRole() != sparkruntime.UserRole_Candidate {
		return nil, fmt.Errorf("only can switch HR or Candidate Role")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := db.UpdateUserRole(ctx, db.DB, req.GetUserRole(), userId); err != nil {
		return nil, err
	}

	return &sparkruntime.SwitchUserRoleResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
