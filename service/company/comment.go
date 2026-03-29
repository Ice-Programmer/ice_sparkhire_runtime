package company

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func BuildReplyComment(ctx context.Context, parentId int64, comment *db.CompanyComment) (*db.CompanyComment, error) {
	// 1. find parent comment
	parentComment, err := db.FindCompanyCommentById(ctx, db.DB, parentId)
	if err != nil {
		return nil, err
	}

	// 2. 判断是否为 reply comment 【如果 root id 为 0 说明为根评论】
	if parentComment.RootId == 0 {
		// not reply comment
		comment.RootId = parentComment.Id
		return comment, nil
	}

	// reply comment
	comment.RootId = parentComment.RootId
	comment.ParentId = parentComment.Id
	comment.ReplyUserId = parentComment.UserId

	return comment, nil
}

func BuildCommentInfoList(ctx context.Context, commentList []*db.CompanyComment) ([]*sparkruntime.CommentInfo, error) {
	// find user info
	userIds := utils.MapStructListDistinct(commentList, func(comment *db.CompanyComment) int64 {
		return comment.UserId
	})

	userMap, err := user.BuildUserBasicInfoMap(ctx, userIds)
	if err != nil {
		return nil, err
	}

	rootCommentList := make([]*sparkruntime.CommentInfo, 0, len(commentList))
	for _, comment := range commentList {
		rootCommentList = append(rootCommentList, &sparkruntime.CommentInfo{
			Id:          comment.Id,
			Content:     comment.Content,
			FavoriteCnt: comment.FavoriteCount,
			CreatorInfo: userMap[comment.UserId],
			CreatedAt:   comment.CreatedAt.Unix(),
		})
	}

	return rootCommentList, nil
}
