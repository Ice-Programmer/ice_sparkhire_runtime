package recruitment

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/elasticsearch"
)

func SyncAllRecruitmentToEs(ctx context.Context, req *sparkruntime.SyncAllRecruitmentToEsRequest) (*sparkruntime.SyncAllRecruitmentToRsResponse, error) {
	if err := elasticsearch.SyncAllRecruitmentToES(ctx); err != nil {
		return nil, err
	}

	return &sparkruntime.SyncAllRecruitmentToRsResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
