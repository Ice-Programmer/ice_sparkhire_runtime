package candidate

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
)

func EditCandidateBasicInfo(ctx context.Context, req *sparkruntime.EditCandidateBasicInfoRequest) (*sparkruntime.EditCandidateBasicInfoResponse, error) {
	if err := validateCandidateBasicInfo(req); err != nil {
		return nil, err
	}

	currentUser, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if sparkruntime.UserRole(currentUser.UserRole) != sparkruntime.UserRole_Candidate {
		return nil, fmt.Errorf("current user is not the candidate role")
	}

	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		if err := db.UpdateCandidateByUserId(ctx, db.DB, currentUser.Id, map[string]interface{}{
			"job_status": req.Status,
		}); err != nil {
			return err
		}

		if err := db.UpdateUserById(ctx, db.DB, currentUser.Id, map[string]interface{}{
			"username":    req.Username,
			"user_avatar": req.Avatar,
			"gender":      req.Gender,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCandidateBasicInfoResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateCandidateBasicInfo(req *sparkruntime.EditCandidateBasicInfoRequest) error {
	if len(req.GetUsername()) == 0 {
		return fmt.Errorf("username can not be empty")
	}

	if len(req.GetUsername()) > 20 {
		return fmt.Errorf("username can not be more than 20 characters")
	}

	if len(req.GetAvatar()) == 0 {
		return fmt.Errorf("avatar can not be empty")
	}

	if req.GetStatus() < sparkruntime.JobStatus_Available || req.GetStatus() > sparkruntime.JobStatus_NotInterested {
		return fmt.Errorf("invalid status")
	}

	return nil
}
