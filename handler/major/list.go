package major

import (
	"context"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/major"
	"ice_sparkhire_runtime/service/redis"
)

func ListMajor(ctx context.Context, req *sparkruntime.ListMajorRequest) (*sparkruntime.ListMajorResponse, error) {

	majorInfoList, err := redis.GetOrSaveList(ctx, consts.MajorListKey, consts.InformationListDuration, major.ListMajor)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ListMajorResponse{
		MajorList: majorInfoList,
		BaseResp:  handler.ConstructSuccessResp(),
	}, nil
}
