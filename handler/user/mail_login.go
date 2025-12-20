package user

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func UserMailLogin(ctx context.Context, req *sparkruntime.UserMailLoginRequest) (*sparkruntime.UserMailLoginResponse, error) {
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

	// 3. get dbUser by email
	dbUser, err := db.FindUserByEmail(ctx, db.DB, req.Email)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, fmt.Errorf("dbUser not found")
	}

	// 4. generate token
	token, err := utils.GenerateToken(consts.TokenExpireTime, dbUser.Id)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.UserMailLoginResponse{
		AccessToken: token,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
