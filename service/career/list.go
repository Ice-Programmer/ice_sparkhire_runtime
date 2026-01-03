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
			CareerTypeName: careerTypeNameMap[careerDB.CareerType],
			CareerTypeId:   careerDB.CareerType,
		}
	})

	return careerInfoList, nil
}
