package information

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/service/redis"
)

func ListGeo(ctx context.Context, req *sparkruntime.ListGeoRequest) (*sparkruntime.ListGeoResponse, error) {
	if req.GetLevel() < sparkruntime.GeoLevel_FirstGeo || req.GetLevel() > sparkruntime.GeoLevel_ForthGeo {
		return nil, fmt.Errorf("geo level out of range")
	}

	if req.GetLevel() != sparkruntime.GeoLevel_FirstGeo && req.GetParentId() <= 0 {
		return nil, fmt.Errorf("parent id is invalid")
	}

	geoInfoList, err := redis.GetOrSaveList(
		ctx,
		fmt.Sprintf("%s:leve_%d:%d", consts.GeoListKey, req.GetLevel(), req.GetParentId()),
		consts.InformationListDuration,
		func(ctx context.Context) ([]*sparkruntime.GeoInfo, error) {
			return geo.ListGeo(ctx, req.GetLevel(), req.GetParentId())
		},
	)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ListGeoResponse{
		GeoList:  geoInfoList,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
