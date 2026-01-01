package candidate

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
)

func EditCandidateProfile(ctx context.Context, req *sparkruntime.EditCandidateProfileRequest) (*sparkruntime.EditCandidateProfileResponse, error) {
	// fetch current user
	currentUser, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if sparkruntime.UserRole(currentUser.UserRole) != sparkruntime.UserRole_Candidate {
		return nil, fmt.Errorf("current user is not the candidate role")
	}

	if err = db.UpdateCandidateByUserId(ctx, db.DB, currentUser.Id, map[string]interface{}{
		"profile": req.GetProfile(),
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCandidateProfileResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
