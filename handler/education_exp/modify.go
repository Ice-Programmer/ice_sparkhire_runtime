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

func ModifyEducationExp(ctx context.Context, req *sparkruntime.ModifyEducationExpRequest) (resp *sparkruntime.ModifyEducationExpResponse, err error) {
	if err = validateModifyEducationExp(ctx, req); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	modifyEducationExp := &db.EducationExperience{
		UserId:    userId,
		SchoolId:  req.SchoolId,
		BeginYear: req.GetBeginYear(),
		EndYear:   req.EndYear,
		MajorId:   req.MajorId,
		Activity:  req.GetActivity(),
	}

	if req.IsSetId() {
		// 编辑
		modifyEducationExp.Id = req.GetId()
		if err := db.UpdateEducationExperience(ctx, db.DB, req.GetId(), modifyEducationExp); err != nil {
			return nil, err
		}
	} else {
		// 新增
		modifyEducationExp.Id = utils.GetId()
		modifyEducationExp.EducationStatus = int8(req.GetStatus())
		if err = db.CreateEducationExperience(ctx, db.DB, modifyEducationExp); err != nil {
			return nil, err
		}

	}

	return &sparkruntime.ModifyEducationExpResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateModifyEducationExp(ctx context.Context, req *sparkruntime.ModifyEducationExpRequest) error {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return err
	}

	// 判断是否修改
	if req.GetId() > 0 {
		// 新增逻辑
		educationExp, err := db.FindEducationExperienceById(ctx, db.DB, req.GetId())
		if err != nil {
			return fmt.Errorf("education exp not found by id %v", req.GetId())
		}

		if educationExp.UserId != userId {
			return fmt.Errorf("only can edit own education experience")
		}
	} else {
		// 判断 status
		if !utils.Contains(sparkruntime.EducationStatusList, req.GetStatus()) {
			return fmt.Errorf("education status is empty")
		}

		// 新增逻辑
		education, err := db.FindEducationExperienceByUserIdAndStatus(ctx, db.DB, userId, req.GetStatus())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("education exp not found by id %v", req.GetId())
		}
		if education != nil {
			return fmt.Errorf("education exp status %s has existed", req.GetStatus().String())
		}
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
