package information

import (
	"context"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/industry"
	"ice_sparkhire_runtime/service/redis"
)

func ListIndustry(ctx context.Context, req *sparkruntime.ListIndustryRequest) (*sparkruntime.ListIndustryResponse, error) {

	industryInfoList, err := redis.GetOrSaveList(ctx, consts.IndustryListKey, consts.InformationListDuration, industry.ListIndustry)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ListIndustryResponse{
		IndustryInfoList: industryInfoList,
		BaseResp:         handler.ConstructSuccessResp(),
	}, nil
}
