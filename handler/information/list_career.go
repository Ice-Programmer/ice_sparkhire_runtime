package information

import (
	"context"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/career"
	"ice_sparkhire_runtime/service/redis"
)

func ListCareer(ctx context.Context, req *sparkruntime.ListCareerInfoRequest) (*sparkruntime.ListCareerInfoResponse, error) {
	careerList, err := redis.GetOrSaveList(ctx, consts.CareerListKey, consts.InformationListDuration, career.ListCareer)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ListCareerInfoResponse{
		CareerList: careerList,
		BaseResp:   handler.ConstructSuccessResp(),
	}, nil
}
