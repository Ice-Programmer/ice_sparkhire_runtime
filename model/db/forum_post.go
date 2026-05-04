package db

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type ForumPost struct {
	Id            int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	UserId        int64          `gorm:"column:user_id;type:bigint;comment:创建用户 id;not null;" json:"user_id"`                                // 创建用户 id
	Title         string         `gorm:"column:title;type:varchar(256);comment:帖子标题;not null;" json:"title"`                                 // 帖子标题
	Content       string         `gorm:"column:content;type:text;comment:内容;" json:"content"`                                                // 内容
	FavoriteCount int32          `gorm:"column:favorite_count;type:int;comment:收藏次数;not null;default:0;" json:"favorite_count"`              // 收藏次数
	ViewCount     int64          `gorm:"column:view_count;type:bigint;comment:浏览次数;not null;default:0;" json:"view_count"`                   // 浏览次数
	Status        int8           `gorm:"column:status;type:tinyint;comment:状态：1-正常 2-审核中 3-屏蔽;not null;default:1;" json:"status"`            // 状态：1-正常 2-审核中 3-屏蔽
	Type          int8           `gorm:"column:type;type:tinyint;comment:帖子类型：1-普通 2-置顶 3-精华;not null;default:1;" json:"type"`               // 帖子类型：1-普通 2-置顶 3-精华
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (ForumPost) TableName() string {
	return "forum_post"
}

func CreateForumPost(ctx context.Context, db *gorm.DB, forumPost *ForumPost) error {
	if err := db.WithContext(ctx).Model(&ForumPost{}).Create(forumPost).Error; err != nil {
		klog.CtxErrorf(ctx, "[ForumPost DB] create post err: %+v", err)
		return err
	}

	return nil
}

func FetchForumPost(ctx context.Context, db *gorm.DB, id int64) (*ForumPost, error) {
	var forumPost ForumPost
	err := db.WithContext(ctx).Model(&ForumPost{}).
		Where("id = ?", id).
		First(&forumPost).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[ForumPost DB] fetch post err: %+v", err)
		return nil, fmt.Errorf("fetch post err: %+v", err)
	}
	return &forumPost, nil
}

func IncrForumPostViewCount(ctx context.Context, db *gorm.DB, postId int64) error {
	err := db.WithContext(ctx).Model(&ForumPost{}).
		Where("id = ?", postId).
		Update("view_count", gorm.Expr("view_count + ?", 1)).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[ForumPost DB] update view_count err: %+v", err)
		return fmt.Errorf("update view_count err: %+v", err)
	}
	return nil
}

func QueryForumPostPage(ctx context.Context, db *gorm.DB, pageSize, pageNum int32, condition *sparkruntime.ForumPostCondition) ([]*ForumPost, int64, error) {
	query := buildForumPostCondition(db, condition)

	var total int64
	if err := query.WithContext(ctx).Count(&total).Error; err != nil {
		klog.CtxErrorf(ctx, "[ForumPost DB] query err: %+v", err)
		return nil, 0, err
	}

	offset := int((pageNum - 1) * pageSize)
	var postList []*ForumPost
	err := query.WithContext(ctx).
		Offset(offset).
		Limit(int(pageSize)).
		Order("created_at DESC").
		Find(&postList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[ForumPost DB] query err: %+v", err)
		return nil, 0, err
	}
	return postList, total, nil
}

func buildForumPostCondition(db *gorm.DB, condition *sparkruntime.ForumPostCondition) *gorm.DB {
	db = db.Model(&ForumPost{})

	if condition == nil {
		return db
	}

	if condition.IsSetSearchText() {
		db = db.Where("title LIKE ?", condition.SearchText).
			Or("content LIKE ? ", condition.SearchText)
	}

	return db
}
