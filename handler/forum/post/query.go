package post

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func QueryForumPostPage(ctx context.Context, req *sparkruntime.QueryForumPostPageRequest) (*sparkruntime.QueryForumPostPageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	postList, total, err := db.QueryForumPostPage(ctx, db.DB, pageSize, pageNum, req.GetCondition())
	if err != nil {
		return nil, err
	}

	postInfoList, err := buildForumPostInfoList(ctx, postList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.QueryForumPostPageResponse{
		Total:    total,
		PostList: postInfoList,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func buildForumPostInfoList(ctx context.Context, postList []*db.ForumPost) ([]*sparkruntime.ForumPostInfo, error) {
	if len(postList) == 0 {
		return nil, nil
	}

	userIds := utils.MapStructListDistinct(postList, func(forumPost *db.ForumPost) int64 {
		return forumPost.UserId
	})

	userBasicInfoMap, err := user.GetUserBasicInfoMap(ctx, userIds)
	if err != nil {
		return nil, err
	}

	return utils.MapStructList(postList, func(post *db.ForumPost) *sparkruntime.ForumPostInfo {
		return &sparkruntime.ForumPostInfo{
			Id:             post.Id,
			Title:          post.Title,
			Content:        post.Content,
			FavouriteCount: post.FavoriteCount,
			ViewCount:      post.ViewCount,
			Status:         sparkruntime.PostStatus(post.Status),
			Type:           sparkruntime.PostType(post.Type),
			CreatedAt:      post.CreatedAt.Unix(),
			CreatorInfo:    userBasicInfoMap[post.UserId],
		}
	}), nil
}
