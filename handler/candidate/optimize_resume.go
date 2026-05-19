package candidate

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino/schema"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/model"
	"ice_sparkhire_runtime/service/profile"
)

func SmartOptimizeCandidateResume(ctx context.Context, req *sparkruntime.SmartOptimizeCandidateResumeRequest) (*sparkruntime.SmartOptimizeCandidateResumeResponse, error) {
	candidateProfileMap, err := profile.GetCandidateProfileMap(ctx)
	if err != nil {
		return nil, err
	}
	// ask AI
	prompt, err := model.GetResumeOptimizePrompt()
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

	var optimizeResumeResult sparkruntime.OptimizeResumeResult_
	if err := sonic.UnmarshalString(outMsg.Content, &optimizeResumeResult); err != nil {
		return nil, err
	}

	return &sparkruntime.SmartOptimizeCandidateResumeResponse{
		OptimizeResumeResult_: &optimizeResumeResult,
		BaseResp:              handler.ConstructSuccessResp(),
	}, nil
}
