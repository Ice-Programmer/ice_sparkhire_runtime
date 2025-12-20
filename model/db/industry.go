package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type Industry struct {
	Id           int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	IndustryName string    `gorm:"column:industry_name;type:varchar(256);comment:行业名称;not null;" json:"industry_name"`                 // 行业名称
	IndustryType int64     `gorm:"column:industry_type;type:bigint;comment:行业类型;not null;" json:"industry_type"`                       // 行业类型
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (Industry) TableName() string {
	return "industry"
}

func ListAllIndustry(ctx context.Context, db *gorm.DB) ([]*Industry, error) {
	var industryList []*Industry
	if err := db.WithContext(ctx).Find(&industryList).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] list all industry error: %v", err)
		return nil, err
	}
	return industryList, nil
}

func FindIndustryById(ctx context.Context, db *gorm.DB, id int64) (*Industry, error) {
	var industry Industry
	err := db.WithContext(ctx).Model(&Industry{}).
		Where("id = ?", id).
		First(&industry).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find industry by id %v error: %v", id, err)
		return nil, err
	}
	return &industry, nil
}
