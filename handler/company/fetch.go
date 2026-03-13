package company

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/company"
)

func FetchCompanyDetailInfo(ctx context.Context, req *sparkruntime.FetchCompanyDetailInfoRequest) (*sparkruntime.FetchCompanyDetailInfoResponse, error) {
	if req.GetCompanyId() <= 0 {
		return nil, fmt.Errorf("invalid company id")
	}

	companyDB, err := db.FindCompanyById(ctx, db.DB, req.CompanyId)
	if err != nil {
		return nil, err
	}

	companyInfo, err := company.BuildCompanyInfo(ctx, companyDB)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.FetchCompanyDetailInfoResponse{
		CompanyInfo: companyInfo,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
