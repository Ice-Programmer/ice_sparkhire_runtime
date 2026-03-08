package company

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/service/industry"
)

func BuildCompanyInfo(ctx context.Context, company *db.Company) (*sparkruntime.CompanyInfo, error) {
	industryDB, err := db.FindIndustryById(ctx, db.DB, company.IndustryId)
	if err != nil {
		return nil, err
	}

	geoDetailInfo, err := geo.BuildGeoDetailInfo(ctx,
		company.FirstGeoLevelId,
		company.SecondGeoLevelId,
		company.ThirdGeoLevelId,
		company.ForthGeoLevelId,
		company.Address,
		company.Longitude,
		company.Latitude,
	)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.CompanyInfo{
		Id:           company.ID,
		CompanyName:  company.CompanyName,
		CompanyLogo:  company.Logo,
		Description:  company.Description,
		GeoInfo:      geoDetailInfo,
		IndustryInfo: industry.BuildIndustryDetail(industryDB),
	}, nil
}
