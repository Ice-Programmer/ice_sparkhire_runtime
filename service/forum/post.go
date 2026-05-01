package forum

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
)

func BuildForumPostInfo(ctx context.Context, post *db.ForumPost) (*sparkruntime.ForumPostInfo, error) {
	// fetch creator info
	creatorInfo, err := user.GetUserBasicInfo(ctx, post.UserId)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.ForumPostInfo{
		Id:             post.Id,
		Title:          post.Title,
		Content:        post.Content,
		FavouriteCount: post.FavoriteCount,
		ViewCount:      post.ViewCount,
		Status:         sparkruntime.ForumStatus(post.Status),
		Type:           sparkruntime.ForumType(post.Type),
		CreatedAt:      post.CreatedAt.Unix(),
		CreatorInfo:    creatorInfo,
	}, nil
}
