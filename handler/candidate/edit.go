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

func EditCandidate(ctx context.Context, req *sparkruntime.EditCandidateRequest) (*sparkruntime.EditCandidateResponse, error) {
	candidateInfo := req.GetCandidateInfo()
	if err := candidate.ValidateCandidate(candidateInfo); err != nil {
		return nil, err
	}

	currentUserId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	originalCandidate, err := db.FindCandidateByUserId(ctx, db.DB, currentUserId)
	if err != nil {
		return nil, err
	}
	if originalCandidate == nil {
		return nil, fmt.Errorf("candidate not found")
	}

	candidateDB, err := candidate.BuildCandidateDB(ctx, req.GetCandidateInfo(), originalCandidate.Id)
	if err != nil {
		return nil, err
	}

	if err = db.UpdateCandidate(ctx, db.DB, candidateDB); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCandidateResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
