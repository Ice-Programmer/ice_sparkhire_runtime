package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type Major struct {
	Id        int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	MajorName string    `gorm:"column:major_name;type:varchar(256);comment:专业名称;not null;" json:"major_name"`                       // 专业名称
	PostNum   int32     `gorm:"column:post_num;type:int;comment:相关数量;not null;default:0;" json:"post_num"`                          // 相关数量
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
}

func (Major) TableName() string {
	return "major"
}

func ListAllMajor(ctx context.Context, db *gorm.DB) ([]*Major, error) {
	var majorList []*Major
	err := db.WithContext(ctx).Find(&majorList).Error
	if err != nil {
		return nil, err
	}
	return majorList, nil
}

func FindMajorById(ctx context.Context, db *gorm.DB, id int64) (*Major, error) {
	var major Major
	err := db.WithContext(ctx).Model(&Major{}).
		Where("id = ?", id).
		First(&major).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find Major by id err: %v", err)
		return nil, err
	}

	return &major, nil
}

func FindMajorListByIds(ctx context.Context, db *gorm.DB, ids []int64) ([]*Major, error) {
	var majorList []*Major
	err := db.WithContext(ctx).Model(&Major{}).
		Where("id IN (?)", ids).
		Find(&majorList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find MajorList err: %v", err)
		return nil, err
	}
	return majorList, nil
}
