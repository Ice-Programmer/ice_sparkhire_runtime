package education_exp

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

func CreateEducationExp(ctx context.Context, req *sparkruntime.CreateEducationExpRequest) (resp *sparkruntime.CreateEducationExpResponse, err error) {
	if err = validateCreateEducationExp(ctx, req); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	oldEducationExp, err := db.FindEducationExperienceByUserIdAndStatus(ctx, db.DB, userId, req.GetStatus())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if oldEducationExp != nil {
		return nil, fmt.Errorf("already have %v education exp", req.GetStatus().String())
	}

	educationExperience := &db.EducationExperience{
		Id:              utils.GetId(),
		UserId:          userId,
		SchoolId:        req.SchoolId,
		EducationStatus: int8(req.Status),
		BeginYear:       req.GetBeginYear(),
		EndYear:         req.EndYear,
		MajorId:         req.MajorId,
		Activity:        req.GetActivity(),
	}

	if err = db.CreateEducationExperience(ctx, db.DB, educationExperience); err != nil {
		return nil, err
	}

	return &sparkruntime.CreateEducationExpResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateCreateEducationExp(ctx context.Context, req *sparkruntime.CreateEducationExpRequest) error {
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
