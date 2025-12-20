package industry

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListIndustry(ctx context.Context) ([]*sparkruntime.IndustryInfo, error) {
	industryTypeList, err := db.ListAllIndustryType(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	industryTypeMap := utils.ToMap(industryTypeList,
		func(industryType *db.IndustryType) int64 { return industryType.Id },
		func(industryType *db.IndustryType) string { return industryType.Name },
	)

	industryListFromDB, err := db.ListAllIndustry(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	industryMap := utils.GroupBy(industryListFromDB,
		func(industry *db.Industry) int64 { return industry.IndustryType },
		func(industry *db.Industry) *db.Industry { return industry },
	)

	industryInfoList := make([]*sparkruntime.IndustryInfo, 0, len(industryTypeList))
	for industryType, industryList := range industryMap {
		if _, ok := industryTypeMap[industryType]; !ok {
			continue
		}

		industryDetailList := utils.MapStructList(industryList, func(industry *db.Industry) *sparkruntime.IndustryDetail {
			return &sparkruntime.IndustryDetail{
				Id:           industry.Id,
				IndustryName: industry.IndustryName,
			}
		})

		industryInfoList = append(industryInfoList, &sparkruntime.IndustryInfo{
			IndustryTypeName: industryTypeMap[industryType],
			IndustryList:     industryDetailList,
		})
	}

	return industryInfoList, nil
}
