package chat

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func ListCurrentUserChatSession(ctx context.Context, req *sparkruntime.ListCurrentUserChatSessionRequest) (*sparkruntime.ListCurrentUserChatSessionResponse, error) {
	// get current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	sessionList, err := db.ListChatSessionByCandidateId(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}

	// get company info
	companyNameMap, err := getCompanyNameMap(ctx, sessionList)
	if err != nil {
		return nil, err
	}

	// get recruitment map
	recruitmentNameMap, err := getRecruitmentNameMap(ctx, sessionList)
	if err != nil {
		return nil, err
	}

	userIdList := utils.MapStructList(sessionList, func(session *db.ChatSession) int64 {
		return session.HrUserId
	})

	userBasicInfoMap, err := user.GetUserBasicInfoMap(ctx, userIdList)
	if err != nil {
		return nil, err
	}

	chatSessionInfos := utils.MapStructList(sessionList, func(session *db.ChatSession) *sparkruntime.ChatSessionInfo {
		return &sparkruntime.ChatSessionInfo{
			Id:                     session.Id,
			CompanyId:              session.CompanyId,
			CompanyName:            companyNameMap[session.CompanyId],
			RecruitmentName:        recruitmentNameMap[session.RecruitmentId],
			RecruitmentId:          session.RecruitmentId,
			ReceiverInfo:           userBasicInfoMap[session.HrUserId],
			LatestMessageCreatedAt: session.LastMessageTime.Unix(),
			UnreadNum:              session.CandidateUnread,
			LastMessage:            session.LastMessage,
		}
	})

	return &sparkruntime.ListCurrentUserChatSessionResponse{
		SessionList: chatSessionInfos,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}

func getCompanyNameMap(ctx context.Context, sessionList []*db.ChatSession) (map[int64]string, error) {
	companyIdList := utils.MapStructList(sessionList, func(session *db.ChatSession) int64 {
		return session.CompanyId
	})

	companyList, err := db.ListCompanyByIds(ctx, db.DB, companyIdList)
	if err != nil {
		return nil, err
	}

	companyNameMap := utils.ToMap(companyList,
		func(company *db.Company) int64 { return company.ID },
		func(company *db.Company) string { return company.CompanyName },
	)

	return companyNameMap, nil
}

func getRecruitmentNameMap(ctx context.Context, sessionList []*db.ChatSession) (map[int64]string, error) {
	recruitmentIdList := utils.MapStructList(sessionList, func(session *db.ChatSession) int64 {
		return session.RecruitmentId
	})

	recruitmentList, err := db.ListRecruitmentByIds(ctx, db.DB, recruitmentIdList)
	if err != nil {
		return nil, err
	}

	recruitmentNameMap := utils.ToMap(recruitmentList,
		func(company *db.Recruitment) int64 { return company.ID },
		func(company *db.Recruitment) string { return company.Name },
	)

	return recruitmentNameMap, nil
}
