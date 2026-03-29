package company

import (
	"context"
	"ice_sparkhire_runtime/model/db"
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
