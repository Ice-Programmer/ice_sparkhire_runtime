package user

import (
	"context"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func RegisterUser(ctx context.Context, req *sparkruntime.RegisterUserRequest) (*sparkruntime.RegisterUserResponse, error) {
	// 1. validate param
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}

	if err := utils.ValidateVerifyCode(req.GetVerifyCode()); err != nil {
		return nil, err
	}

	// 2. compare verify code
	if err := user.CompareVerifyCode(ctx, req.Email, req.VerifyCode); err != nil {
		return nil, err
	}

	// 4. gen empty user
	emptyUser, err := user.BuildEmptyUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// 5. generate token
	token, err := utils.GenerateToken(consts.TokenExpireTime, emptyUser.Id)
	if err != nil {
		return nil, err
	}

	// 6. save db
	if err = db.SaveUser(ctx, db.DB, emptyUser); err != nil {
		return nil, err
	}

	return &sparkruntime.RegisterUserResponse{
		AccessToken: token,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
