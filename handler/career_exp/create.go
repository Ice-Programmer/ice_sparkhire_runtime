package career_exp

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func CreateCareerExperience(ctx context.Context, req *sparkruntime.CreateCareerExperienceRequest) (*sparkruntime.CreateCareerExperienceResponse, error) {
	if err := validateCreateCareerExperience(req); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] get current user id failed, %v", err)
		return nil, err
	}

	if err = db.CreateCareerExperience(ctx, db.DB, &db.CareerExperience{
		Id:             utils.GetId(),
		UserId:         userId,
		ExperienceName: req.ExperienceName,
		StartTS:        req.StartTS,
		EndTS:          req.EndTS,
		JobRole:        req.JobRole,
		Description:    req.Description,
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.CreateCareerExperienceResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateCreateCareerExperience(req *sparkruntime.CreateCareerExperienceRequest) error {
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
