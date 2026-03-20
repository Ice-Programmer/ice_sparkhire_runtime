package user

import (
	"context"
	"fmt"
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

	// target is exist
	var err error
	switch req.GetTargetType() {
	case sparkruntime.TargetType_Recruitment:
		_, err = db.FindRecruitmentById(ctx, db.DB, req.GetTargetId())
	case sparkruntime.TargetType_Company:
		_, err = db.FindCompanyById(ctx, db.DB, req.GetTargetId())
	}
	if err != nil {
		return nil, fmt.Errorf("find %s error", req.GetTargetType().String())
	}

	// fetch current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current user id error: %w", err)
	}

	// upsert favor record (里面处理了冲突)
	err = db.UpsertUserFavor(ctx, db.DB, &db.UserFavorite{
		Id:         utils.GetId(),
		UserId:     userId,
		TargetType: int8(req.GetTargetType()),
		TargetId:   req.GetTargetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("favor %s error: %w", req.GetTargetType().String(), err)
	}

	return &sparkruntime.UserFavorResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
