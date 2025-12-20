package company

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/tos"
	"ice_sparkhire_runtime/utils"
)

func CreateCompany(ctx context.Context, req *sparkruntime.CreateCompanyRequest) (*sparkruntime.CreateCompanyResponse, error) {
	if err := validateCreateCompany(ctx, req); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	avatarBytes, err := utils.GenerateCharAvatar('陈', 300)
	if err != nil {
		return nil, err
	}

	logoFileName := fmt.Sprintf("%s.png", req.GetCompanyName())
	fileKey := tos.BuildFileKey(logoFileName, sparkruntime.FileType_CompanyAvatar)

	if err = tos.UploadObject(ctx, fileKey, avatarBytes); err != nil {
		return nil, err
	}

	company := &db.Company{
		ID:           utils.GetId(),
		CompanyName:  req.CompanyName,
		CreateUserId: userId,
		Description:  req.Description,
		Logo:         tos.UrlPrefix + fileKey,
		IndustryId:   req.GetIndustryId(),
	}

	if req.IsSetGeoInfo() {
		company.FirstGeoLevelId = req.GetGeoInfo().FirstGeoLevelId
		company.SecondGeoLevelId = req.GetGeoInfo().SecondGeoLevelId
		company.ThirdGeoLevelId = req.GetGeoInfo().ThirdGeoLevelId
		company.ForthGeoLevelId = req.GetGeoInfo().ForthGeoLevelId
		company.Longitude = req.GetGeoInfo().Longitude
		company.Address = req.GetGeoInfo().Address
		company.Latitude = req.GetGeoInfo().Latitude
	}

	if err = db.CreateCompany(ctx, db.DB, company); err != nil {
		return nil, err
	}

	return &sparkruntime.CreateCompanyResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateCreateCompany(ctx context.Context, req *sparkruntime.CreateCompanyRequest) error {
	if len(req.GetCompanyName()) == 0 {
		return fmt.Errorf("company name is required")
	}

	if len(req.GetCompanyName()) > 50 {
		return fmt.Errorf("company name is too long")
	}

	if len(req.GetDescription()) == 0 {
		return fmt.Errorf("company description is required")
	}

	if len(req.GetDescription()) > 2000 {
		return fmt.Errorf("company description is too long")
	}

	if req.GetIndustryId() <= 0 {
		return fmt.Errorf("industry id is invalid")
	}

	if _, err := db.FindIndustryById(ctx, db.DB, req.GetIndustryId()); err != nil {
		return fmt.Errorf("find industry error: %v", err)
	}

	company, err := db.FindCompanyByName(ctx, db.DB, req.GetCompanyName())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if company != nil {
		return fmt.Errorf("company %s already exists", req.GetCompanyName())
	}

	return nil
}
