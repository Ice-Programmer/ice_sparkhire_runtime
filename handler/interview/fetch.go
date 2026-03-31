package interview

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/interview"
	"ice_sparkhire_runtime/utils"
)

func FetchCurrentUserInterview(ctx context.Context, req *sparkruntime.FetchCurrentUserInterviewRequest) (*sparkruntime.FetchCurrentUserInterviewResponse, error) {
	// 1. get current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	interviewScheduleList, err := db.ListInterviewScheduleByCandidateId(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}

	interviewInfoList, err := interview.BuildInterviewInfoList(ctx, interviewScheduleList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.FetchCurrentUserInterviewResponse{
		InterviewList: interviewInfoList,
		BaseResp:      handler.ConstructSuccessResp(),
	}, nil
}
