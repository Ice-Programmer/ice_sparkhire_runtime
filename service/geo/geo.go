package geo

import (
	"context"
	"fmt"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListGeo(ctx context.Context, level sparkruntime.GeoLevel, parentId int64) ([]*sparkruntime.GeoInfo, error) {
	switch level {
	case sparkruntime.GeoLevel_FirstGeo:
		return ListFirstGeo(ctx)
	case sparkruntime.GeoLevel_SecondGeo:
		return ListSecondGeo(ctx, parentId)
	case sparkruntime.GeoLevel_ThirdGeo:
		return ListThirdGeo(ctx, parentId)
	case sparkruntime.GeoLevel_ForthGeo:
		return ListForthGeo(ctx, parentId)
	}
	return nil, fmt.Errorf("invalid geo level %v", level)
}

func ListFirstGeo(ctx context.Context) ([]*sparkruntime.GeoInfo, error) {
	geoFirstLevelList, err := db.ListAllFirstGeo(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	return utils.MapStructList(geoFirstLevelList, func(geoFirstLevel *db.GeoFirstLevel) *sparkruntime.GeoInfo {
		return &sparkruntime.GeoInfo{
			Id:    geoFirstLevel.Id,
			Name:  geoFirstLevel.GeoName,
			Code:  geoFirstLevel.Code,
			Level: sparkruntime.GeoLevel_FirstGeo,
		}
	}), nil
}

func ListSecondGeo(ctx context.Context, firstGeoId int64) ([]*sparkruntime.GeoInfo, error) {
	geoList, err := db.ListSecondGeoByFirstGeoId(ctx, db.DB, firstGeoId)
	if err != nil {
		return nil, err
	}

	return utils.MapStructList(geoList, func(geoFirstLevel *db.GeoSecondLevel) *sparkruntime.GeoInfo {
		return &sparkruntime.GeoInfo{
			Id:    geoFirstLevel.Id,
			Name:  geoFirstLevel.GeoName,
			Code:  geoFirstLevel.Code,
			Level: sparkruntime.GeoLevel_SecondGeo,
		}
	}), nil
}

func ListThirdGeo(ctx context.Context, secondGeoId int64) ([]*sparkruntime.GeoInfo, error) {
	geoList, err := db.ListThirdGeoBySecondGeoId(ctx, db.DB, secondGeoId)
	if err != nil {
		return nil, err
	}

	return utils.MapStructList(geoList, func(geoFirstLevel *db.GeoThirdLevel) *sparkruntime.GeoInfo {
		return &sparkruntime.GeoInfo{
			Id:    geoFirstLevel.Id,
			Name:  geoFirstLevel.GeoName,
			Code:  geoFirstLevel.Code,
			Level: sparkruntime.GeoLevel_ThirdGeo,
		}
	}), nil
}

func ListForthGeo(ctx context.Context, thirdGeoId int64) ([]*sparkruntime.GeoInfo, error) {
	geoList, err := db.ListForthGeoByThirdGeoId(ctx, db.DB, thirdGeoId)
	if err != nil {
		return nil, err
	}

	return utils.MapStructList(geoList, func(geoFirstLevel *db.GeoForthLevel) *sparkruntime.GeoInfo {
		return &sparkruntime.GeoInfo{
			Id:    geoFirstLevel.Id,
			Name:  geoFirstLevel.GeoName,
			Code:  geoFirstLevel.Code,
			Level: sparkruntime.GeoLevel_ForthGeo,
		}
	}), nil
}

func BuildGeoDetailInfo(ctx context.Context, firstGeoLevelId, secondGeoLevelId, thirdGeoLevelId, forthGeoLevelId int64, address string) (*sparkruntime.GeoDetailInfo, error) {
	firstGeo, err := db.FindFirstGeoById(ctx, db.DB, firstGeoLevelId)
	if err != nil {
		return nil, err
	}

	secondGeo, err := db.FindSecondGeoById(ctx, db.DB, secondGeoLevelId)
	if err != nil {
		return nil, err
	}

	thirdGeo, err := db.FindThirdLevelById(ctx, db.DB, thirdGeoLevelId)
	if err != nil {
		return nil, err
	}

	forthGeo, err := db.FindForthLevelById(ctx, db.DB, forthGeoLevelId)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.GeoDetailInfo{
		FirstGeoLevelId:    firstGeoLevelId,
		FirstGeoLevelName:  utils.StringPtr(firstGeo.GeoName),
		SecondGeoLevelId:   secondGeoLevelId,
		SecondGeoLevelName: utils.StringPtr(secondGeo.GeoName),
		ThirdGeoLevelId:    thirdGeoLevelId,
		ThirdGeoLevelName:  utils.StringPtr(thirdGeo.GeoName),
		ForthGeoLevelId:    forthGeoLevelId,
		ForthGeoLevelName:  utils.StringPtr(forthGeo.GeoName),
		Address:            utils.StringPtr(address),
	}, nil
}
