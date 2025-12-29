package company

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func EditCompany(ctx context.Context, req *sparkruntime.EditCompanyRequest) (*sparkruntime.EditCompanyResponse, error) {
	if err := validateEditCompany(ctx, req); err != nil {
		return nil, err
	}

	company, err := db.FindCompanyById(ctx, db.DB, req.Id)
	if err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if company.CreateUserId != userId {
		return nil, fmt.Errorf("only creator can edit company")
	}

	companyLogo := company.Logo
	if req.GetLogo() != company.Logo {
		companyLogo = req.GetLogo()
	}

	modifyCompany := &db.Company{
		ID:          req.Id,
		CompanyName: req.CompanyName,
		Description: req.Description,
		Logo:        companyLogo,
		IndustryId:  req.GetIndustryId(),
	}

	if req.IsSetGeoInfo() {
		modifyCompany.FirstGeoLevelId = req.GetGeoInfo().FirstGeoLevelId
		modifyCompany.SecondGeoLevelId = req.GetGeoInfo().SecondGeoLevelId
		modifyCompany.ThirdGeoLevelId = req.GetGeoInfo().ThirdGeoLevelId
		modifyCompany.ForthGeoLevelId = req.GetGeoInfo().ForthGeoLevelId
		modifyCompany.Address = req.GetGeoInfo().Address
		modifyCompany.Longitude = req.GetGeoInfo().Longitude
		modifyCompany.Latitude = req.GetGeoInfo().Latitude
	}

	if err = db.UpdateCompany(ctx, db.DB, req.Id, modifyCompany); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCompanyResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateEditCompany(ctx context.Context, req *sparkruntime.EditCompanyRequest) error {
	if req.GetId() <= 0 {
		return fmt.Errorf("invalid company id")
	}

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

	return nil
}
