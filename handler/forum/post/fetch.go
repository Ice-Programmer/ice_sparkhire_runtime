package post

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/forum"
)

func FetchForumPost(ctx context.Context, req *sparkruntime.FetchForumPostRequest) (*sparkruntime.FetchForumPostResponse, error) {
	if req.GetId() <= 0 {
		return nil, errors.New("invalid id")
	}

	forumPost, err := db.FetchForumPost(ctx, db.DB, req.Id)
	if err != nil {
		return nil, err
	}

	forumPostInfo, err := forum.BuildForumPostInfo(ctx, forumPost)
	if err != nil {
		return nil, err
	}

	// incr view count
	// todo 改造 redis
	if err := db.IncrForumPostViewCount(ctx, db.DB, forumPost.Id); err != nil {
		klog.CtxErrorf(ctx, "IncrForumPostViewCount error: %v", err)
	}

	return &sparkruntime.FetchForumPostResponse{
		PostInfo: forumPostInfo,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
