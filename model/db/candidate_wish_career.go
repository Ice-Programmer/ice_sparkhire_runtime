package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const CandidateWishCareerTableName = "candidate_wish_career"

type CandidateWishCareer struct {
	Id            int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	UserId        int64          `gorm:"column:user_id;type:bigint;comment:用户id;not null;" json:"user_id"`                                   // 用户id
	CareerId      int64          `gorm:"column:career_id;type:bigint;comment:职业id;not null;default:0;" json:"career_id"`                     // 职业id
	SalaryUpper   int32          `gorm:"column:salary_upper;type:int;comment:薪水上限;not null;default:0;" json:"salary_upper"`                  // 薪水上限
	SalaryLower   int32          `gorm:"column:salary_lower;type:int;comment:薪水下限;not null;default:0;" json:"salary_lower"`                  // 薪水下限
	CurrencyType  int32          `gorm:"column:currency_type;type:int;comment:薪水货币类型;not null;default:1;" json:"currency_type"`              // 薪水货币类型
	FrequencyType int8           `gorm:"column:frequency_type;type:tinyint;comment:类型;not null;default:1;" json:"frequency_type"`            // 类型
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (CandidateWishCareer) TableName() string {
	return CandidateWishCareerTableName
}

func CreateWishCareer(ctx context.Context, db *gorm.DB, wishCareer *CandidateWishCareer) error {
	if err := db.WithContext(ctx).Create(wishCareer).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] create wish career fail, %v", err)
		return err
	}
	return nil
}

func FindWishCareerByUserIdAndCareerId(ctx context.Context, db *gorm.DB, userId int64, careerId int64) (*CandidateWishCareer, error) {
	var wishCareer CandidateWishCareer
	err := db.WithContext(ctx).Model(&CandidateWishCareer{}).
		Where("user_id = ?", userId).
		Where("career_id = ?", careerId).
		First(&wishCareer).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find wish career by user id fail, %v", err)
		return nil, err
	}

	return &wishCareer, nil
}

func FindWishCareerById(ctx context.Context, db *gorm.DB, id int64) (*CandidateWishCareer, error) {
	var wishCareer CandidateWishCareer
	err := db.WithContext(ctx).Model(&CandidateWishCareer{}).
		Where("id = ?", id).
		First(&wishCareer).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find wish career by id fail, %v", err)
		return nil, err
	}
	return &wishCareer, nil
}

func UpdateWishCareerById(ctx context.Context, db *gorm.DB, id int64, updateMap map[string]interface{}) error {
	err := db.WithContext(ctx).Model(&CandidateWishCareer{}).
		Where("id = ?", id).
		Updates(updateMap).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update wish career by id fail, %v", err)
		return err
	}
	return nil
}

func FindWishCareerByUserId(ctx context.Context, db *gorm.DB, userId int64) ([]*CandidateWishCareer, error) {
	var wishCareers []*CandidateWishCareer
	err := db.WithContext(ctx).Model(&CandidateWishCareer{}).
		Where("user_id = ?", userId).
		Find(&wishCareers).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find wish careers fail, %v", err)
		return nil, err
	}
	return wishCareers, nil
}

func DeleteWishCareerById(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Model(&CandidateWishCareer{}).
		Where("id = ?", id).
		Delete(&CandidateWishCareer{}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] delete wish career by id fail, %v", err)
		return err
	}
	return nil
}
