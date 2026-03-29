package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type CompanyComment struct {
	Id            int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CompanyId     int64          `gorm:"column:company_id;type:bigint;comment:公司 id;not null;" json:"company_id"`                            // 公司 id
	UserId        int64          `gorm:"column:user_id;type:bigint;comment:用户 id;not null;" json:"user_id"`                                  // 用户 id
	Content       string         `gorm:"column:content;type:text;comment:评论内容;" json:"content"`                                              // 评论内容
	RootId        int64          `gorm:"column:root_id;type:bigint;comment:所属根评论 id，0 表示自身为根;not null;default:0;" json:"root_id"`            // 所属根评论 id，0 表示自身为根
	ParentId      int64          `gorm:"column:parent_id;type:bigint;comment:被回复帖子，0 为根节点;not null;default:0;" json:"parent_id"`             // 被回复帖子，0 为根节点
	ReplyUserId   int64          `gorm:"column:reply_user_id;type:bigint;comment:被回复用户 id;not null;default:0;" json:"reply_user_id"`         // 被回复用户 id
	ReplyCount    int32          `gorm:"column:reply_count;type:int;comment:若是根评论，记录其子评论总数;not null;default:0;" json:"reply_count"`          // 若是根评论，记录其子评论总数
	FavoriteCount int32          `gorm:"column:favorite_count;type:int;comment:点赞数量;not null;default:0;" json:"favorite_count"`              // 点赞数量
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (*CompanyComment) TableName() string {
	return "company_comment"
}

func CreateCompanyComment(ctx context.Context, db *gorm.DB, comment *CompanyComment) error {
	if err := db.WithContext(ctx).Model(&CompanyComment{}).Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func FindCompanyCommentById(ctx context.Context, db *gorm.DB, id int64) (*CompanyComment, error) {
	var comment CompanyComment
	err := db.WithContext(ctx).Model(&CompanyComment{}).
		Where("id = ?", id).
		First(&comment).Error

	if err != nil {
		klog.Errorf("[CompanyComment] Find by id error: %v", err)
		return nil, err
	}
	return &comment, nil
}

func IncrCompanyCommentReplyCount(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Model(&CompanyComment{}).
		Where("id = ?", id).
		Update("reply_count", gorm.Expr("reply_count + ?", 1)).Error
	if err != nil {
		klog.Errorf("[CompanyComment] Incr comment reply count error: %v", err)
		return err
	}

	return nil
}
