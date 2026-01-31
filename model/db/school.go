package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type School struct {
	Id         int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	SchoolName string    `gorm:"column:school_name;type:varchar(256);comment:学校名称;not null;" json:"school_name"`                     // 学校名称
	PostNum    int64     `gorm:"column:post_num;type:bigint;comment:相关数量;not null;default:0;" json:"post_num"`                       // 相关数量
	SchoolIcon string    `gorm:"column:school_icon;type:varchar(256);comment:学校icon;not null;" json:"school_icon"`                   // 学校 icon
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (School) TableName() string {
	return "school"
}

func ListAllSchool(ctx context.Context, db *gorm.DB) ([]*School, error) {
	var schoolList []*School
	if err := db.WithContext(ctx).Find(&schoolList).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] list all school error: %v", err)
		return nil, err
	}
	return schoolList, nil
}

func FindSchoolById(ctx context.Context, db *gorm.DB, id int64) (*School, error) {
	var school School
	err := db.WithContext(ctx).Model(&School{}).
		Where("id = ?", id).
		First(&school).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find school by id %v error: %v", id, err)
		return nil, err
	}
	return &school, nil
}

func FindSchoolListByIds(ctx context.Context, db *gorm.DB, ids []int64) ([]*School, error) {
	var schoolList []*School
	err := db.WithContext(ctx).Model(&School{}).
		Where("id IN (?)", ids).
		Find(&schoolList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find school by ids %v error: %v", ids, err)
		return nil, err
	}
	return schoolList, nil
}
