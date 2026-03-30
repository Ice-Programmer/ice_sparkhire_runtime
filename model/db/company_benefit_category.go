package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type CompanyBenefitCategory struct {
	Id        int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CompanyId int64          `gorm:"column:company_id;type:bigint;comment:公司 id;not null;" json:"company_id"`                            // 公司 id
	Title     string         `gorm:"column:title;type:varchar(128);comment:福利分组标题;not null;" json:"title"`                               // 福利分组标题
	Subtitle  string         `gorm:"column:subtitle;type:varchar(255);comment:分组说明;default:NULL;" json:"subtitle"`                       // 分组说明
	Sort      int32          `gorm:"column:sort;type:int;comment:排序，越小越靠前;not null;default:0;" json:"sort"`                              // 排序，越小越靠前
	Status    int8           `gorm:"column:status;type:tinyint;comment:状态：1正常，-1禁用;not null;default:1;" json:"status"`                   // 状态：1正常，-1禁用
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (CompanyBenefitCategory) TableName() string {
	return "company_benefit_category"
}

func FindCompanyBenefitCategoryByCompanyId(ctx context.Context, db *gorm.DB, companyId int64) ([]*CompanyBenefitCategory, error) {
	var benefitCategoryList []*CompanyBenefitCategory
	err := db.WithContext(ctx).Model(&CompanyBenefitCategory{}).
		Where("company_id = ?", companyId).
		Order("sort").
		Find(&benefitCategoryList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[CompanyBenefitCategory] find category by company Id %d error: %v", companyId, err)
	}

	return benefitCategoryList, err
}
