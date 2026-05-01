package post

import (
	"context"
	"errors"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func CreateForumPost(ctx context.Context, req *sparkruntime.CreateForumPostRequest) (*sparkruntime.CreateForumPostResponse, error) {
	if err := validatePost(req.GetTitle(), req.GetContent()); err != nil {
		return nil, err
	}

	// fetch current user
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if err := db.CreateForumPost(ctx, db.DB, &db.ForumPost{
		Id:      utils.GetId(),
		UserId:  userId,
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}); err != nil {
		return nil, err
	}

	return &sparkruntime.CreateForumPostResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validatePost(title, content string) error {
	if len(title) == 0 {
		return errors.New("post title can not be empty")
	}

	if len(title) > 40 {
		return errors.New("post title can not be more than 40 characters")
	}

	if len(content) > 5000 {
		return errors.New("post content can not be more than 5000 characters")
	}

	return nil
}
