package recruitment

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/career"
	"ice_sparkhire_runtime/service/company"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/utils"
)

func CreateRecruitment(ctx context.Context, req *sparkruntime.CreateRecruitmentRequest) (*sparkruntime.CreateRecruitmentResponse, error) {
	// 1. validate create recruitment
	if err := validateCreateRecruitment(ctx, req); err != nil {
		return nil, err
	}

	// 2. generate recruitment db
	recruitment, err := buildRecruitment(ctx, req)
	if err != nil {
		return nil, err
	}

	// create recruitment
	if err := db.CreateRecruitment(ctx, db.DB, recruitment); err != nil {
		return nil, err
	}

	return &sparkruntime.CreateRecruitmentResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func buildRecruitment(ctx context.Context, req *sparkruntime.CreateRecruitmentRequest) (*db.Recruitment, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	recruitment := &db.Recruitment{
		ID:            utils.GetId(),
		Name:          req.Name,
		UserId:        userId,
		CompanyId:     req.CompanyId,
		CareerId:      req.CareerId,
		Description:   req.Description,
		Requirement:   req.Requirement,
		EducationType: int8(req.EducationStatus),
		JobType:       int8(req.JobType),
	}

	if req.IsSetGeoInfo() {
		geoInfo := req.GetGeoInfo()
		recruitment.FirstGeoLevelId = geoInfo.FirstGeoLevelId
		recruitment.SecondGeoLevelId = geoInfo.SecondGeoLevelId
		recruitment.ThirdGeoLevelId = geoInfo.ThirdGeoLevelId
		recruitment.ForthGeoLevelId = geoInfo.ForthGeoLevelId
		recruitment.Address = geoInfo.Address
		recruitment.Longitude = geoInfo.Longitude
		recruitment.Latitude = geoInfo.Latitude
	}

	if req.IsSetSalaryInfo() {
		salaryInfo := req.GetSalaryInfo()
		recruitment.CurrencyType = int32(salaryInfo.GetCurrencyType())
		recruitment.FrequencyType = int8(salaryInfo.GetFrequencyType())
		recruitment.SalaryUpper = salaryInfo.GetSalaryUpper()
		recruitment.SalaryLower = salaryInfo.GetSalaryLower()
	}

	return recruitment, nil
}

func validateCreateRecruitment(ctx context.Context, req *sparkruntime.CreateRecruitmentRequest) error {
	// 1. company
	if err := company.ValidateCompanyId(ctx, req.GetCompanyId()); err != nil {
		return err
	}

	// 2. career
	if err := career.ValidateCareerId(ctx, req.GetCareerId()); err != nil {
		return err
	}

	// 3. description
	if err := utils.ValidateStrLen(req.Description, 1000); err != nil {
		return fmt.Errorf("description: %s", err)
	}

	// 4. requirement
	if err := utils.ValidateStrLen(req.Requirement, 1000); err != nil {
		return fmt.Errorf("requirement: %s", err)
	}

	// Name
	if err := utils.ValidateStrLen(req.Name, 20); err != nil {
		return fmt.Errorf("recruitment name: %s", err)
	}

	// 5. geo
	if err := geo.ValidateGeoInfo(ctx, req.GetGeoInfo()); err != nil {
		return fmt.Errorf("geo info: %s", err)
	}

	// 6. education status
	if req.GetEducationStatus().String() == consts.EnumNotFound {
		return fmt.Errorf("education status is invalid")
	}

	// 7. job type
	if req.GetJobType().String() == consts.EnumNotFound {
		return fmt.Errorf("job type is invalid")
	}

	// 8. salary
	if err := career.ValidateSalaryInfo(req.GetSalaryInfo()); err != nil {
		return fmt.Errorf("salary info: %s", err)
	}

	return nil
}
