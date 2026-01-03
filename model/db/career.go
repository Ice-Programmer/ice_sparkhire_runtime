package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const CareerTableName = "career"

type Career struct {
	Id          int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CareerName  string         `gorm:"column:career_name;type:varchar(256);comment:职业名称;not null;" json:"career_name"`                     // 职业名称
	Description string         `gorm:"column:description;type:varchar(1024);comment:职业介绍;" json:"description"`                             // 职业介绍
	CareerIcon  string         `gorm:"column:career_icon;type:varchar(256);comment:icon;" json:"career_icon"`                              // icon
	CareerType  int64          `gorm:"column:career_type;type:int;comment:职业类型;not null;" json:"career_type"`                              // 职业类型
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (Career) TableName() string {
	return CareerTableName
}

const CareerTypeTableName = "career_type"

type CareerType struct {
	ID             int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CareerTypeName string         `gorm:"column:career_type_name;type:varchar(256);comment:职业类型名称;not null;" json:"career_type_name"`         // 职业类型名称
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (CareerType) TableName() string {
	return CareerTypeTableName
}

func ListCareer(ctx context.Context, db *gorm.DB) ([]*Career, error) {
	var careers []*Career
	if err := db.WithContext(ctx).Model(&Career{}).Find(&careers).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] list career err: %v", err)
		return nil, err
	}
	return careers, nil
}

func ListCareerType(ctx context.Context, db *gorm.DB) ([]*CareerType, error) {
	var careerTypes []*CareerType
	if err := db.WithContext(ctx).Model(&CareerType{}).Find(&careerTypes).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] list careerType err: %v", err)
		return nil, err
	}
	return careerTypes, nil
}

func FindCareerById(ctx context.Context, db *gorm.DB, id int64) (*Career, error) {
	var career Career
	err := db.WithContext(ctx).Model(&Career{}).
		Where("id = ?", id).
		Find(&career).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find career err: %v", err)
		return nil, err
	}
	return &career, nil
}

func FindCareerByIds(ctx context.Context, db *gorm.DB, ids []int64) ([]*Career, error) {
	var careers []*Career
	err := db.WithContext(ctx).Model(&Career{}).
		Where("id IN (?)", ids).
		Find(&careers).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find careers err: %v", err)
		return nil, err
	}
	return careers, nil
}
