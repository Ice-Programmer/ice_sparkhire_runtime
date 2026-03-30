package company

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/company"
)

func FetchCompanyBenefits(ctx context.Context, req *sparkruntime.FetchCompanyBenefitsRequest) (*sparkruntime.FetchCompanyBenefitsResponse, error) {
	if req.GetCompanyId() <= 0 {
		return nil, fmt.Errorf("invalid company id")
	}

	// 1. exist company
	if _, err := db.FindCompanyById(ctx, db.DB, req.GetCompanyId()); err != nil {
		return nil, fmt.Errorf("failed to fetch company: %w", err)
	}

	// fetch benefit info list
	benefitInfoList, err := company.BuildBenefitInfoList(ctx, req.CompanyId)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.FetchCompanyBenefitsResponse{
		BenefitList: benefitInfoList,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
