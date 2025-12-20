package career_exp

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func GetCurrentUserCareerExperience(ctx context.Context, req *sparkruntime.GetCurrentUserCareerExperienceRequest) (*sparkruntime.GetCurrentUserCareerExperienceResponse, error) {
	// 1. get current login user
	currentUserInfo, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	careerExperienceList, err := db.FindCareerExperienceByUserId(ctx, db.DB, currentUserInfo.Id)
	if err != nil {
		return nil, err
	}

	experienceInfoList := utils.MapStructList(careerExperienceList, func(experience *db.CareerExperience) *sparkruntime.CareerExperienceInfo {
		return &sparkruntime.CareerExperienceInfo{
			Id:             experience.Id,
			JobRole:        experience.JobRole,
			Description:    experience.Description,
			StartTS:        experience.StartTS,
			EndTS:          experience.EndTS,
			ExperienceName: experience.ExperienceName,
		}
	})

	return &sparkruntime.GetCurrentUserCareerExperienceResponse{
		CareerExperienceInfoList: experienceInfoList,
		BaseResp:                 handler.ConstructSuccessResp(),
	}, nil
}
