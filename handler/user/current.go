package user

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func FetchCurrentUser(ctx context.Context, req *sparkruntime.FetchCurrentUserRequest) (*sparkruntime.FetchCurrentUserResponse, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	userInfo, err := db.FindUserById(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.FetchCurrentUserResponse{
		BasicInfo: user.BuildUserBasicInfo(userInfo),
		BaseResp:  handler.ConstructSuccessResp(),
	}, nil
}
