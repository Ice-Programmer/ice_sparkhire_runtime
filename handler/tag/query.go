package tag

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func QueryTag(ctx context.Context, req *sparkruntime.QueryTagRequest) (*sparkruntime.QueryTagResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	tagList, total, err := db.QueryTagPage(ctx, db.DB, pageSize, pageNum, req.GetSearchText())
	if err != nil {
		return nil, err
	}

	tagInfoList := utils.MapStructList(tagList, func(tag *db.Tag) *sparkruntime.TagInfo {
		return &sparkruntime.TagInfo{
			Id:      tag.Id,
			TagName: tag.TagName,
		}
	})

	return &sparkruntime.QueryTagResponse{
		TagList:  tagInfoList,
		Total:    total,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
