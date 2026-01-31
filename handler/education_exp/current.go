package education_exp

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func GetCurrentEducationExp(ctx context.Context, req *sparkruntime.GetCurrentUserEducationExpRequest) (*sparkruntime.GetCurrentUserEducationExpResponse, error) {
	userInfo, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if sparkruntime.UserRole(userInfo.UserRole) != sparkruntime.UserRole_Candidate {
		return nil, fmt.Errorf("user is not candidate")
	}

	educationExperienceList, err := db.FindEducationExpListByUserId(ctx, db.DB, userInfo.Id)
	if err != nil {
		return nil, err
	}

	educationExpInfoList, err := buildEducationExpInfoList(ctx, educationExperienceList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.GetCurrentUserEducationExpResponse{
		EducationExpList: educationExpInfoList,
		BaseResp:         handler.ConstructSuccessResp(),
	}, nil
}

func buildEducationExpInfoList(ctx context.Context, educationExperiences []*db.EducationExperience) ([]*sparkruntime.EducationExpInfo, error) {
	majorIdList := make([]int64, 0, len(educationExperiences))
	schoolIdList := make([]int64, 0, len(educationExperiences))
	for _, educationExperience := range educationExperiences {
		majorIdList = append(majorIdList, educationExperience.MajorId)
		schoolIdList = append(schoolIdList, educationExperience.SchoolId)
	}

	majorList, err := db.FindMajorListByIds(ctx, db.DB, majorIdList)
	if err != nil {
		return nil, err
	}

	majorMap := utils.ToMap(majorList,
		func(major *db.Major) int64 { return major.Id },
		func(major *db.Major) string { return major.MajorName },
	)

	schoolList, err := db.FindSchoolListByIds(ctx, db.DB, schoolIdList)
	if err != nil {
		return nil, err
	}

	schoolMap := utils.ToMap(schoolList,
		func(school *db.School) int64 { return school.Id },
		func(school *db.School) *db.School { return school },
	)

	educationExpInfoList := make([]*sparkruntime.EducationExpInfo, 0, len(educationExperiences))
	for _, educationExperience := range educationExperiences {
		educationExpInfoList = append(educationExpInfoList, &sparkruntime.EducationExpInfo{
			Id:        educationExperience.Id,
			BeginYear: educationExperience.BeginYear,
			EndYear:   educationExperience.EndYear,
			Activity:  educationExperience.Activity,
			MajorInfo: &sparkruntime.MajorInfo{
				Id:        educationExperience.MajorId,
				MajorName: majorMap[educationExperience.MajorId],
			},
			SchoolInfo: &sparkruntime.SchoolInfo{
				Id:         educationExperience.SchoolId,
				SchoolName: schoolMap[educationExperience.SchoolId].SchoolName,
				SchoolIcon: schoolMap[educationExperience.SchoolId].SchoolIcon,
			},
		})
	}
	return educationExpInfoList, nil
}
