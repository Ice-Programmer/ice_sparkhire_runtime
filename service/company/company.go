package company

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/service/industry"
	"ice_sparkhire_runtime/utils"
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

func BuildCompanyInfoMapByIds(ctx context.Context, companyIds []int64) (map[int64]*sparkruntime.CompanyInfo, error) {
	companyList, err := db.FindCompanyByIds(ctx, db.DB, companyIds)
	if err != nil {
		return nil, err
	}

	industryIdList := utils.MapStructListDistinct(companyList, func(company *db.Company) int64 {
		return company.IndustryId
	})

	industryMap, err := industry.BuildIndustryDetailMapByIds(ctx, industryIdList)
	if err != nil {
		return nil, err
	}

	companyInfoMap := utils.ToMap(companyList,
		func(company *db.Company) int64 { return company.ID },
		func(company *db.Company) *sparkruntime.CompanyInfo {
			geoFullNameInfo, err := db.FindGeoFullInfo(ctx, db.DB, company.ForthGeoLevelId)
			if err != nil {
				return nil
			}

			return &sparkruntime.CompanyInfo{
				Id:          company.ID,
				CompanyName: company.CompanyName,
				CompanyLogo: company.Logo,
				Description: company.Description,
				GeoInfo: &sparkruntime.GeoDetailInfo{
					FirstGeoLevelId:    company.FirstGeoLevelId,
					SecondGeoLevelId:   company.SecondGeoLevelId,
					ThirdGeoLevelId:    company.ThirdGeoLevelId,
					ForthGeoLevelId:    company.ForthGeoLevelId,
					Address:            company.Address,
					Latitude:           company.Latitude,
					Longitude:          company.Longitude,
					FirstGeoLevelName:  geoFullNameInfo.FirstName,
					SecondGeoLevelName: geoFullNameInfo.SecondName,
					ThirdGeoLevelName:  geoFullNameInfo.ThirdName,
					ForthGeoLevelName:  geoFullNameInfo.ForthName,
				},
				IndustryInfo: industryMap[company.IndustryId],
			}
		},
	)

	return companyInfoMap, nil
}
