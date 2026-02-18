package tag

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func GetCurrentCandidateTags(ctx context.Context, req *sparkruntime.GetCurrentCandidateTagsRequest) (*sparkruntime.GetCurrentCandidateTagsResponse, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	tagRels, err := db.FindTagRelsByObjId(ctx, db.DB, userId, int32(sparkruntime.TagObjType_Candidate))
	if err != nil {
		return nil, err
	}

	tagIdList := utils.MapStructList(tagRels, func(tagRel *db.TagObjRel) int64 {
		return tagRel.TagId
	})

	tagList, err := db.FindTagByIds(ctx, db.DB, tagIdList)
	if err != nil {
		return nil, err
	}

	tagInfoList := utils.MapStructList(tagList, func(tag *db.Tag) *sparkruntime.TagInfo {
		return &sparkruntime.TagInfo{
			Id:      tag.Id,
			TagName: tag.TagName,
		}
	})

	return &sparkruntime.GetCurrentCandidateTagsResponse{
		TagList:  tagInfoList,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
