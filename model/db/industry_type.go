package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type IndustryType struct {
	Id        int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	Name      string    `gorm:"column:name;type:varchar(256);comment:行业类型名称;not null;" json:"name"`                                 // 行业类型名称
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (IndustryType) TableName() string {
	return "industry_type"
}

func ListAllIndustryType(ctx context.Context, db *gorm.DB) ([]*IndustryType, error) {
	var industryTypeList []*IndustryType
	if err := db.WithContext(ctx).Find(&industryTypeList).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] list all industry type error: %v", err)
		return nil, err
	}
	return industryTypeList, nil
}
