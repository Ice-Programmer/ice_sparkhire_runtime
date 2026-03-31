package tag

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/model/neo4j"
	"ice_sparkhire_runtime/utils"
)

func AddTag(ctx context.Context, req *sparkruntime.AddTagRequest) (*sparkruntime.AddTagResponse, error) {
	if err := validateTag(req.GetTagName()); err != nil {
		return nil, err
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

	tagId := utils.GetId()

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		_, err := db.AddTag(ctx, tx, &db.Tag{
			Id:           tagId,
			TagName:      req.GetTagName(),
			CreateUserId: userId,
		})
		if err != nil {
			return err
		}

		if err := neo4j.InsertTagNode(ctx, &neo4j.TagNode{
			ID:      tagId,
			TagName: req.TagName,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.AddTagResponse{
		Id:       tagId,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateTag(tagName string) error {
	if len(tagName) == 0 {
		return fmt.Errorf("tag name is required")
	}

	if len(tagName) > 200 {
		return fmt.Errorf("tag name is too long")
	}

	return nil
}
