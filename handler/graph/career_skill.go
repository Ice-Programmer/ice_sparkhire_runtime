package graph

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/neo4j"
	"ice_sparkhire_runtime/utils"
)

func FetchCareerRelativeSkillTags(ctx context.Context, req *sparkruntime.FetchCareerRelativeSkillTagsRequest) (*sparkruntime.FetchCareerRelativeSkillTagsResponse, error) {
	if req.GetCareerId() <= 0 {
		return nil, fmt.Errorf("invalid career_id")
	}

	tags, err := neo4j.FindCareerRelativeTags(ctx, req.CareerId)
	if err != nil {
		return nil, err
	}

	tagList := utils.MapStructList(tags, func(tagNode *neo4j.TagNode) *sparkruntime.TagInfo {
		return &sparkruntime.TagInfo{
			Id:      tagNode.ID,
			TagName: tagNode.TagName,
		}
	})

	return &sparkruntime.FetchCareerRelativeSkillTagsResponse{
		TagList:  tagList,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
