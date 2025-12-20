package tag

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func AddTag(ctx context.Context, req *sparkruntime.AddTagRequest) (*sparkruntime.AddTagResponse, error) {
	if len(req.GetTagName()) == 0 {
		return nil, fmt.Errorf("tag name is required")
	}

	if len(req.GetTagName()) > 20 {
		return nil, fmt.Errorf("tag name is too long")
	}

	tag, err := db.FindTagByName(ctx, db.DB, req.GetTagName())
	if err != nil {
		return nil, err
	}
	if tag != nil {
		return nil, fmt.Errorf("tag already exists")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	tagId, err := db.AddTag(ctx, db.DB, &db.Tag{
		Id:           utils.GetId(),
		TagName:      req.GetTagName(),
		CreateUserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &sparkruntime.AddTagResponse{
		Id:       tagId,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
