package career

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListCareer(ctx context.Context) ([]*sparkruntime.CareerInfo, error) {
	careerTypeList, err := db.ListCareerType(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	careerTypeNameMap := utils.ToMap(careerTypeList,
		func(careerType *db.CareerType) int64 { return careerType.ID },
		func(careerType *db.CareerType) string { return careerType.CareerTypeName },
	)

	careerList, err := db.ListCareer(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	careerInfoList := utils.MapStructList(careerList, func(careerDB *db.Career) *sparkruntime.CareerInfo {
		return &sparkruntime.CareerInfo{
			Id:             careerDB.Id,
			CareerName:     careerDB.CareerName,
			CareerTypeName: utils.StringPtr(careerTypeNameMap[careerDB.CareerType]),
			CareerTypeId:   utils.Int64Ptr(careerDB.CareerType),
			CareerIcon:     careerDB.CareerIcon,
			Description:    careerDB.Description,
		}
	})

	return careerInfoList, nil
}

func BuildCareerMapByIds(ctx context.Context, careerIds []int64) (map[int64]*sparkruntime.CareerInfo, error) {
	careerList, err := db.FindCareerByIds(ctx, db.DB, careerIds)
	if err != nil {
		return nil, err
	}

	careerTypeIdList := utils.MapStructList(careerList, func(career *db.Career) int64 {
		return career.CareerType
	})

	careerTypeList, err := db.FindCareerTypeByIds(ctx, db.DB, careerTypeIdList)
	if err != nil {
		return nil, err
	}

	careerTypeNameMap := utils.ToMap(careerTypeList,
		func(careerType *db.CareerType) int64 { return careerType.ID },
		func(careerType *db.CareerType) string { return careerType.CareerTypeName },
	)

	careerInfoMap := utils.ToMap(careerList,
		func(career *db.Career) int64 { return career.Id },
		func(career *db.Career) *sparkruntime.CareerInfo {
			return &sparkruntime.CareerInfo{
				Id:             career.Id,
				CareerName:     career.CareerName,
				CareerTypeName: utils.StringPtr(careerTypeNameMap[career.CareerType]),
				CareerTypeId:   utils.Int64Ptr(career.CareerType),
				CareerIcon:     career.CareerIcon,
				Description:    career.Description,
			}
		},
	)

	return careerInfoMap, nil
}
