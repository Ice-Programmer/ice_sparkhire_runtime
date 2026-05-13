package candidate

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino/schema"
	"ice_sparkhire_runtime/handler"
	"ice_sparkhire_runtime/handler/career_exp"
	"ice_sparkhire_runtime/handler/education_exp"
	"ice_sparkhire_runtime/handler/tag"
	"ice_sparkhire_runtime/handler/wish_career"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/model"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func SmartOptimizeCandidateResume(ctx context.Context, req *sparkruntime.SmartOptimizeCandidateResumeRequest) (*sparkruntime.SmartOptimizeCandidateResumeResponse, error) {
	// 1. get user basic info
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}
	userBasicInfo, err := user.GetUserBasicInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 2. education exp list
	educationExpResp, err := education_exp.GetCurrentEducationExp(ctx, &sparkruntime.GetCurrentUserEducationExpRequest{})
	if err != nil {
		return nil, err
	}

	// 3. career exp list
	careerExpResp, err := career_exp.GetCurrentUserCareerExperience(ctx, &sparkruntime.GetCurrentUserCareerExperienceRequest{})
	if err != nil {
		return nil, err
	}

	// tag list
	tagsResp, err := tag.GetCurrentCandidateTags(ctx, &sparkruntime.GetCurrentCandidateTagsRequest{})
	if err != nil {
		return nil, err
	}

	// 4. wish exp list
	wishCareerResp, err := wish_career.GetCurrentWishCareer(ctx, &sparkruntime.GetCurrentWishCareerRequest{})
	if err != nil {
		return nil, err
	}

	// 5. ask AI
	prompt, err := model.GetResumeOptimizePrompt()
	if err != nil {
		return nil, err
	}

	renderPrompt, err := model.RenderPrompt(prompt, map[string]string{
		"USER_BASIC_INFO":    utils.MarshalString(userBasicInfo),
		"EDUCATION_EXP_LIST": utils.MarshalString(educationExpResp.GetEducationExpList()),
		"CAREER_EXP_LIST":    utils.MarshalString(careerExpResp.GetCareerExperienceInfoList()),
		"TAG_LIST":           utils.MarshalString(tagsResp.GetTagList()),
		"WISH_CAREER_LIST":   utils.MarshalString(wishCareerResp.GetWishCareerList()),
	})

	chatModel, err := model.NewCommonChatModel(ctx)
	if err != nil {
		return nil, err
	}

	outMsg, err := chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: renderPrompt,
		},
	})
	if err != nil {
		return nil, err
	}

	var optimizeResumeResult sparkruntime.OptimizeResumeResult_
	if err := sonic.UnmarshalString(outMsg.Content, &optimizeResumeResult); err != nil {
		return nil, err
	}

	return &sparkruntime.SmartOptimizeCandidateResumeResponse{
		OptimizeResumeResult_: &optimizeResumeResult,
		BaseResp:              handler.ConstructSuccessResp(),
	}, nil
}
