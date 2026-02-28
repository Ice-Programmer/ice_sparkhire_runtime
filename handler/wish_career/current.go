package wish_career

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func GetCurrentWishCareer(ctx context.Context, req *sparkruntime.GetCurrentWishCareerRequest) (*sparkruntime.GetCurrentWishCareerResponse, error) {
	currentUserId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	wishCareerList, err := db.FindWishCareerByUserId(ctx, db.DB, currentUserId)
	if err != nil {
		return nil, err
	}

	wishCareerInfoList, err := buildWishCareerInfoList(ctx, wishCareerList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.GetCurrentWishCareerResponse{
		WishCareerList: wishCareerInfoList,
		BaseResp:       handler.ConstructSuccessResp(),
	}, nil
}

func buildWishCareerInfoList(ctx context.Context, wishCareerList []*db.CandidateWishCareer) ([]*sparkruntime.WishCareerInfo, error) {
	careerIdList := utils.MapStructList(wishCareerList, func(wishCareer *db.CandidateWishCareer) int64 {
		return wishCareer.CareerId
	})

	careerList, err := db.FindCareerByIds(ctx, db.DB, careerIdList)
	if err != nil {
		return nil, err
	}

	careerMap := utils.ToMap(careerList,
		func(career *db.Career) int64 { return career.Id },
		func(career *db.Career) *db.Career { return career },
	)

	careerTypeIdSet := utils.MapStructListDistinct(careerList, func(career *db.Career) int64 {
		return career.CareerType
	})

	careerTypeList, err := db.FindCareerTypeByIds(ctx, db.DB, careerTypeIdSet)
	if err != nil {
		return nil, err
	}

	careerTypeNameMap := utils.ToMap(careerTypeList,
		func(careerType *db.CareerType) int64 { return careerType.ID },
		func(careerType *db.CareerType) string { return careerType.CareerTypeName },
	)

	wishCareerInfoList := make([]*sparkruntime.WishCareerInfo, 0, len(wishCareerList))
	for _, wishCareer := range wishCareerList {
		career := careerMap[wishCareer.CareerId]
		wishCareerInfoList = append(wishCareerInfoList, &sparkruntime.WishCareerInfo{
			Id: wishCareer.Id,
			CareerInfo: &sparkruntime.CareerInfo{
				Id:             wishCareer.CareerId,
				CareerName:     career.CareerName,
				CareerIcon:     career.CareerIcon,
				CareerTypeName: utils.StringPtr(careerTypeNameMap[career.CareerType]),
				CareerTypeId:   utils.Int64Ptr(career.CareerType),
			},
			SalaryLower:   wishCareer.SalaryLower,
			SalaryUpper:   wishCareer.SalaryUpper,
			CurrencyType:  sparkruntime.SalaryCurrencyType(wishCareer.CurrencyType),
			FrequencyType: sparkruntime.SalaryFrequencyType(wishCareer.FrequencyType),
		})
	}

	return wishCareerInfoList, nil
}
