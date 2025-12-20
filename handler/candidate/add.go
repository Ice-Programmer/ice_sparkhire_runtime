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

func AddCandidate(ctx context.Context, req *sparkruntime.AddCandidateRequest) (*sparkruntime.AddCandidateResponse, error) {
	candidateInfo := req.GetCandidateInfo()
	if err := candidate.ValidateCandidate(candidateInfo); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	candidateDb, err := db.FindCandidateByUserId(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}
	if candidateDb != nil {
		return nil, fmt.Errorf("candidate already exists")
	}

	candidateDB, err := candidate.BuildCandidateDB(ctx, candidateInfo, utils.GetId())
	if err != nil {
		return nil, err
	}

	if err = db.AddCandidate(ctx, db.DB, candidateDB); err != nil {
		return nil, err
	}

	return &sparkruntime.AddCandidateResponse{
		Id:       candidateDB.Id,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
