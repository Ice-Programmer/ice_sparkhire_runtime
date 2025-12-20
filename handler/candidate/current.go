package candidate

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/candidate"
	"ice_sparkhire_runtime/utils"
)

func GetCurrentCandidate(ctx context.Context, req *sparkruntime.GetCurrentCandidateRequest) (*sparkruntime.GetCurrentCandidateResponse, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	originalCandidate, err := db.FindCandidateByUserId(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}
	if originalCandidate == nil {
		return nil, fmt.Errorf("candidate not found")
	}

	candidateInfo, err := candidate.BuildCandidateInfo(ctx, originalCandidate)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.GetCurrentCandidateResponse{
		CandidateInfo: candidateInfo,
		BaseResp:      handler.ConstructSuccessResp(),
	}, nil
}
