package career_exp

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func EditCareerExperience(ctx context.Context, req *sparkruntime.EditCareerExperienceRequest) (*sparkruntime.EditCareerExperienceResponse, error) {
	if err := validateEditCareerExperience(ctx, req); err != nil {
		return nil, err
	}

	modifyCareerExperience := &db.CareerExperience{
		ExperienceName: req.ExperienceName,
		StartTS:        req.StartTS,
		EndTS:          req.EndTS,
		JobRole:        req.JobRole,
		Description:    req.Description,
	}

	if err := db.UpdateCareerExperience(ctx, db.DB, req.Id, modifyCareerExperience); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCareerExperienceResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateEditCareerExperience(ctx context.Context, req *sparkruntime.EditCareerExperienceRequest) error {
	if req.GetId() <= 0 {
		return fmt.Errorf("invalid id")
	}

	careerExperience, err := db.FindCareerExperienceById(ctx, db.DB, req.GetId())
	if err != nil {
		return fmt.Errorf("find career experience error: %s", err.Error())
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return fmt.Errorf("get current user id error: %s", err.Error())
	}

	if careerExperience.UserId != userId {
		return fmt.Errorf("only can edit own career experience")
	}

	if len(req.ExperienceName) == 0 {
		return fmt.Errorf("career experience name is required")
	}

	if len(req.JobRole) == 0 {
		return fmt.Errorf("career experience job role is required")
	}

	if len(req.ExperienceName) == 0 {
		return fmt.Errorf("career experience name is required")
	}

	if req.GetEndTS() > 0 && req.GetStartTS() > req.GetEndTS() {
		return fmt.Errorf("end ts is elearier than start ts")
	}

	return nil
}
