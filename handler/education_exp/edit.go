package education_exp

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func EditEducationExp(ctx context.Context, req *sparkruntime.EditEducationExpRequest) (*sparkruntime.EditEducationExpResponse, error) {
	if err := validateEditEducationExp(ctx, req); err != nil {
		return nil, err
	}

	modifyEducationExperience := &db.EducationExperience{
		SchoolId:  req.SchoolId,
		BeginYear: req.BeginYear,
		EndYear:   req.EndYear,
		MajorId:   req.MajorId,
		Activity:  req.Activity,
	}

	if err := db.UpdateEducationExperience(ctx, db.DB, req.GetId(), modifyEducationExperience); err != nil {
		return nil, err
	}

	return &sparkruntime.EditEducationExpResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateEditEducationExp(ctx context.Context, req *sparkruntime.EditEducationExpRequest) error {
	if req.GetId() <= 0 {
		return fmt.Errorf("id is invalid")
	}

	educationExp, err := db.FindEducationExperienceById(ctx, db.DB, req.GetId())
	if err != nil {
		return fmt.Errorf("find EducationExperience error: %s", err.Error())
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return err
	}

	if educationExp.UserId != userId {
		return fmt.Errorf("only can edit own education experience")
	}

	if len(req.GetActivity()) > 2000 {
		return fmt.Errorf("activity length can not longer than 2000")
	}

	if err := utils.ValidateBeginYearAndEndYear(req.GetBeginYear(), req.GetEndYear()); err != nil {
		return err
	}

	if req.GetMajorId() <= 0 {
		return fmt.Errorf("major id is invalid")
	}

	if _, err := db.FindMajorById(ctx, db.DB, req.GetMajorId()); err != nil {
		return fmt.Errorf("find major err: %v", err)
	}

	if _, err := db.FindSchoolById(ctx, db.DB, req.GetSchoolId()); err != nil {
		return fmt.Errorf("find school err: %v", err)
	}

	return nil
}
