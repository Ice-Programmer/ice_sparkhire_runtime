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
	"strings"
)

func BatchCreateTag(ctx context.Context, req *sparkruntime.BatchCreateTagRequest) (*sparkruntime.BatchCreateTagResponse, error) {
	if len(req.GetTagNameList()) == 0 {
		return nil, fmt.Errorf("tag list is empty")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	tagList := make([]*db.Tag, 0, len(req.GetTagNameList()))
	tagNodeList := make([]*neo4j.TagNode, 0, len(req.GetTagNameList()))
	for _, tagName := range req.TagNameList {
		if err := validateTag(tagName); err != nil {
			return nil, fmt.Errorf("tag 「%s」 is invalid: %v", tagName, err)
		}

		tagId := utils.GetId()
		tagList = append(tagList, &db.Tag{
			Id:           tagId,
			TagName:      strings.TrimSpace(tagName),
			CreateUserId: userId,
		})

		tagNodeList = append(tagNodeList, &neo4j.TagNode{
			ID:      tagId,
			TagName: strings.ToLower(strings.TrimSpace(tagName)),
		})
	}

	var num int32
	if err := db.DB.Transaction(func(tx *gorm.DB) (err error) {
		// 1. batch add db
		num, err = db.BatchAddTag(ctx, tx, tagList)
		if err != nil {
			return err
		}

		// 2. insert neo4j
		if err := neo4j.BatchInsertTagNode(ctx, tagNodeList); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.BatchCreateTagResponse{
		Total:    num,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
