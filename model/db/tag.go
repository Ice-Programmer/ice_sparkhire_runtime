package db

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/utils"
	"time"
)

const (
	TagTableName = "tag"
)

type Tag struct {
	Id           int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	TagName      string    `gorm:"column:tag_name;type:varchar(256);comment:标签名称;not null;" json:"tag_name"`                           // 标签名称
	CreateUserId int64     `gorm:"column:create_user_id;type:bigint;comment:创建用户 id;not null;" json:"create_user_id"`                  // 创建用户 id
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (Tag) TableName() string {
	return TagTableName
}

func QueryTagPage(ctx context.Context, db *gorm.DB, pageSize, pageNum int32, searchText string) ([]*Tag, int64, error) {
	db = db.WithContext(ctx)

	if len(searchText) > 0 {
		db = db.Where("tag_name LIKE ?", utils.WrapLike(searchText))
	}

	var count int64
	if err := db.Model(Tag{}).Count(&count).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] count tag error: %v", err)
		return nil, 0, err
	}

	var tagList []*Tag
	offset := int((pageNum - 1) * pageSize)
	if err := db.Model(Tag{}).Limit(int(pageSize)).Offset(offset).Find(&tagList).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] query tag error: %v", err)
		return nil, 0, err
	}

	return tagList, count, nil
}

func AddTag(ctx context.Context, db *gorm.DB, tag *Tag) (int64, error) {
	err := db.WithContext(ctx).Save(tag).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] add tag error: %v", err)
		return 0, err
	}
	return tag.Id, nil
}

func FindTagByName(ctx context.Context, db *gorm.DB, name string) (*Tag, error) {
	var tag Tag
	err := db.WithContext(ctx).First(&tag, "tag_name = ?", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find tag error: %v", err)
		return nil, err
	}
	return &tag, nil
}

func FindTagByIds(ctx context.Context, db *gorm.DB, id []int64) ([]*Tag, error) {
	var tagList []*Tag
	err := db.WithContext(ctx).Find(&tagList, "id in (?)", id).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find tag error: %v", err)
		return nil, err
	}
	return tagList, nil
}
