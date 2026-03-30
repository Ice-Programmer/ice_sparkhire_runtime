package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type CompanyBenefitItem struct {
	Id         int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CategoryId int64          `gorm:"column:category_id;type:bigint;comment:福利分组id;not null;" json:"category_id"`                         // 福利分组id
	Title      string         `gorm:"column:title;type:varchar(128);comment:福利条目标题;not null;" json:"title"`                               // 福利条目标题
	Content    string         `gorm:"column:content;type:varchar(512);comment:福利条目描述;default:NULL;" json:"content"`                       // 福利条目描述
	Sort       int32          `gorm:"column:sort;type:int;comment:排序，越小越靠前;not null;default:0;" json:"sort"`                              // 排序，越小越靠前
	Status     int8           `gorm:"column:status;type:tinyint;comment:状态：1正常，-1禁用;not null;default:1;" json:"status"`                   // 状态：1正常，-1禁用
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (CompanyBenefitItem) TableName() string {
	return "company_benefit_item"
}

func FindCompanyBenefitItemByCategoryIds(ctx context.Context, db *gorm.DB, categoryIds []int64) ([]*CompanyBenefitItem, error) {
	var benefitItemList []*CompanyBenefitItem
	err := db.WithContext(ctx).Model(&CompanyBenefitItem{}).
		Where("category_id in (?)", categoryIds).
		Order("sort").
		Find(&benefitItemList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[CompanyBenefitItem] find benefit item list error: %v", err)
		return nil, err
	}
	return benefitItemList, nil
}
