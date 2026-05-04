package interview

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func BuildInterviewInfoList(ctx context.Context, interviewList []*db.InterviewSchedule) ([]*sparkruntime.InterviewInfo, error) {
	if len(interviewList) == 0 {
		return []*sparkruntime.InterviewInfo{}, nil
	}

	// get user info
	userIds := make([]int64, 0, len(interviewList)*2)
	for _, interview := range interviewList {
		userIds = append(userIds, interview.CandidateId, interview.CreatorId)
	}
	userBasicInfoMap, err := user.GetUserBasicInfoMap(ctx, userIds)
	if err != nil {
		return nil, err
	}

	// get company info
	companyIds := utils.MapStructListDistinct(interviewList, func(interview *db.InterviewSchedule) int64 {
		return interview.CompanyId
	})

	companyList, err := db.ListCompanyByIds(ctx, db.DB, companyIds)
	if err != nil {
		return nil, err
	}
	companyMap := utils.ToMap(companyList,
		func(company *db.Company) int64 { return company.ID },
		func(company *db.Company) *db.Company { return company },
	)

	// get recruitment info
	recruitmentIds := utils.MapStructListDistinct(interviewList, func(interview *db.InterviewSchedule) int64 {
		return interview.RecruitmentId
	})
	recruitmentList, err := db.ListRecruitmentByIds(ctx, db.DB, recruitmentIds)
	if err != nil {
		return nil, err
	}
	recruitmentNameMap := utils.ToMap(recruitmentList,
		func(recruitment *db.Recruitment) int64 { return recruitment.ID },
		func(recruitment *db.Recruitment) string { return recruitment.Name },
	)

	interviewInfoList := utils.MapStructList(interviewList, func(interview *db.InterviewSchedule) *sparkruntime.InterviewInfo {
		return &sparkruntime.InterviewInfo{
			Id:              interview.Id,
			CandidateInfo:   userBasicInfoMap[interview.CandidateId],
			CreatorInfo:     userBasicInfoMap[interview.CreatorId],
			RecruitmentId:   interview.RecruitmentId,
			RecruitmentName: recruitmentNameMap[interview.RecruitmentId],
			CompanyId:       interview.CompanyId,
			CompanyName:     companyMap[interview.CompanyId].CompanyName,
			CompanyLink:     companyMap[interview.CompanyId].Logo,
			InterviewTs:     interview.InterviewTs,
			Type:            interview.InterviewType,
			Status:          interview.Status,
			Duration:        interview.InterviewDuration,
			InterviewLink:   interview.InterviewLink,
			InterviewDate:   interview.InterviewDate,
		}
	})

	return interviewInfoList, nil
}
