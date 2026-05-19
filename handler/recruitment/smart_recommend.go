package recruitment

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/model"
	"ice_sparkhire_runtime/service/profile"
	"ice_sparkhire_runtime/service/recruitment"
	"ice_sparkhire_runtime/utils"
)

func SmartRecommendRecruitment(ctx context.Context, req *sparkruntime.SmartRecommendRecruitmentRequest) (*sparkruntime.SmartRecommendRecruitmentResponse, error) {
	candidateProfileMap, err := profile.GetCandidateProfileMap(ctx)
	if err != nil {
		return nil, err
	}

	// 获取 recruitment list
	recruitmentMilvusItems, err := recruitment.SearchFromEmbedding(ctx, utils.MarshalString(candidateProfileMap), 20)
	if err != nil {
		klog.CtxErrorf(ctx, "recruitment SearchFromEmbedding err: %v", err)
		return nil, err
	}
	ids := utils.MapStructList(recruitmentMilvusItems, func(recruitmentMilvus *recruitment.RecruitmentMilvusItem) int64 {
		return recruitmentMilvus.RecruitmentID
	})
	recruitmentList, err := db.ListRecruitmentByIds(ctx, db.DB, ids)
	if err != nil {
		klog.CtxErrorf(ctx, "ListRecruitmentByIds err: %v", err)
		return nil, err
	}
	candidateProfileMap["RECRUITMENT_LIST"] = utils.MarshalString(recruitmentList)

	// 询问 ai
	prompt, err := model.GetRecommendRecruitmentPrompt()
	if err != nil {
		return nil, err
	}

	renderPrompt, err := model.RenderPrompt(prompt, candidateProfileMap)
	if err != nil {
		return nil, err
	}

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

	recommendRecruitmentResult := sparkruntime.RecommendRecruitmentResult_{}
	if err := sonic.UnmarshalString(outMsg.Content, &recommendRecruitmentResult); err != nil {
		return nil, err
	}

	return &sparkruntime.SmartRecommendRecruitmentResponse{
		RecommendRecruitmentResult_: &recommendRecruitmentResult,
		BaseResp:                    handler.ConstructSuccessResp(),
	}, nil
}
