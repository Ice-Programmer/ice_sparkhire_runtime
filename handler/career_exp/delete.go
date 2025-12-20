package career_exp

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func DeleteCareerExperience(ctx context.Context, req *sparkruntime.DeleteCareerExperienceRequest) (*sparkruntime.DeleteCareerExperienceResponse, error) {
	if req.GetId() <= 0 {
		return nil, fmt.Errorf("invalid id")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := db.FindCareerExperienceByIdAndUserId(ctx, db.DB, req.GetId(), userId); err != nil {
		return nil, fmt.Errorf("find CareerExperience error: %v", err)
	}

	if err = db.DeleteEducationExperience(ctx, db.DB, req.GetId()); err != nil {
		return nil, err
	}

	return &sparkruntime.DeleteCareerExperienceResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
