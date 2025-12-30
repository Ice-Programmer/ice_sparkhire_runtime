package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
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

	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		if err := db.UpdateUserRole(ctx, db.DB, req.GetUserRole(), userId); err != nil {
			return err
		}

		var err error
		switch req.GetUserRole() {
		case sparkruntime.UserRole_HR:
			err = UpsertHR(ctx, tx, userId)
		case sparkruntime.UserRole_Candidate:
			err = UpsertCandidateInfo(ctx, tx, userId)
		}

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("switch user role err: %v", err)
	}

	return &sparkruntime.SwitchUserRoleResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func UpsertCandidateInfo(ctx context.Context, tx *gorm.DB, userId int64) error {
	candidateInfo, err := db.FindCandidateByUserId(ctx, tx, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if candidateInfo != nil {
		return nil
	}

	if err = db.CreateCandidate(ctx, tx, &db.Candidate{
		Id:     utils.GetId(),
		UserId: userId,
	}); err != nil {
		return err
	}

	return nil
}

func UpsertHR(ctx context.Context, tx *gorm.DB, userId int64) error {
	// todo 待补充
	return nil
}
