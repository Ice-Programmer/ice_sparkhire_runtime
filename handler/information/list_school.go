package information

import (
	"context"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/redis"
	"ice_sparkhire_runtime/service/school"
)

func ListSchool(ctx context.Context, req *sparkruntime.ListSchoolRequest) (*sparkruntime.ListSchoolResponse, error) {
	schoolInfoList, err := redis.GetOrSaveList(ctx, consts.SchoolListKey, consts.InformationListDuration, school.ListSchool)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ListSchoolResponse{
		SchoolList: schoolInfoList,
		BaseResp:   handler.ConstructSuccessResp(),
	}, nil
}
