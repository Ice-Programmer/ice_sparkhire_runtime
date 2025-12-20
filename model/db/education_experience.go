package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

const EducationExperienceTableName = "education_experience"

type EducationExperience struct {
	Id              int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	UserId          int64          `gorm:"column:user_id;type:bigint;comment:用户id;not null;" json:"user_id"`                                   // 用户id
	SchoolId        int64          `gorm:"column:school_id;type:bigint;comment:学校id;not null;" json:"school_id"`                               // 学校id
	EducationStatus int8           `gorm:"column:education_status;type:tinyint;comment:学历类型;not null;" json:"education_status"`                // 学历类型
	BeginYear       int32          `gorm:"column:begin_year;type:int;comment:开始年份;not null;" json:"begin_year"`                                // 开始年份
	EndYear         int32          `gorm:"column:end_year;type:int;comment:结束年份;not null;" json:"end_year"`                                    // 结束年份
	MajorId         int64          `gorm:"column:major_id;type:bigint;comment:专业id;not null;" json:"major_id"`                                 // 专业id
	Activity        string         `gorm:"column:activity;type:text;comment:在校经历;" json:"activity"`                                            // 在校经历
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (EducationExperience) TableName() string {
	return EducationExperienceTableName
}

func CreateEducationExperience(ctx context.Context, db *gorm.DB, educationExp *EducationExperience) error {
	if err := db.WithContext(ctx).Create(educationExp).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] Create EducationExperience error:%v", err)
		return err
	}

	return nil
}

func FindEducationExperienceByUserIdAndStatus(ctx context.Context, db *gorm.DB, userId int64, status sparkruntime.EducationStatus) (*EducationExperience, error) {
	var educationExperience EducationExperience
	err := db.WithContext(ctx).Model(&EducationExperience{}).
		Where("user_id = ? AND education_status = ?", userId, status).
		First(&educationExperience).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find EducationExperience error:%v", err)
		return nil, err
	}
	return &educationExperience, nil
}

func FindEducationExperienceById(ctx context.Context, db *gorm.DB, id int64) (*EducationExperience, error) {
	var educationExperience EducationExperience
	err := db.WithContext(ctx).Model(&EducationExperience{}).
		Where("id = ?", id).
		First(&educationExperience).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find EducationExperience error: %v", err)
		return nil, err
	}
	return &educationExperience, nil
}

func FindEducationExperienceByUserIdAndId(ctx context.Context, db *gorm.DB, userId int64, id int64) (*EducationExperience, error) {
	var educationExperience EducationExperience
	err := db.WithContext(ctx).Model(&EducationExperience{}).
		Where("user_id = ? AND id = ?", userId, id).
		First(&educationExperience).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find EducationExperience error: %v", err)
		return nil, err
	}
	return &educationExperience, nil
}

func UpdateEducationExperience(ctx context.Context, db *gorm.DB, id int64, modifyEducationExp *EducationExperience) error {
	err := db.WithContext(ctx).Model(&EducationExperience{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"school_id":  modifyEducationExp.SchoolId,
			"begin_year": modifyEducationExp.BeginYear,
			"end_year":   modifyEducationExp.EndYear,
			"major_id":   modifyEducationExp.MajorId,
			"activity":   modifyEducationExp.Activity,
		}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update education experience error: %v", err)
		return err
	}
	return nil
}

func DeleteEducationExperience(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&EducationExperience{}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] delete education experience error: %v", err)
		return err
	}
	return nil
}

func FindEducationExpListByUserId(ctx context.Context, db *gorm.DB, userId int64) ([]*EducationExperience, error) {
	var educationExpList []*EducationExperience
	err := db.WithContext(ctx).Model(&EducationExperience{}).
		Where("user_id = ?", userId).
		Find(&educationExpList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] Find EducationExpByUserId error: %v", err)
		return nil, err
	}
	return educationExpList, nil
}
