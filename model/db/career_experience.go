package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const CareerExperienceTableName = "career_experience"

type CareerExperience struct {
	Id             int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	UserId         int64          `gorm:"column:user_id;type:bigint;comment:用户id;not null;" json:"user_id"`                                   // 用户id
	ExperienceName string         `gorm:"column:experience_name;type:varchar(256);comment:经历名称;not null;" json:"experience_name"`             // 经历名称
	StartTS        int64          `gorm:"column:begin_ts;type:bigint;comment:开始时间;not null;default:0;" json:"begin_ts"`                       // 开始时间
	EndTS          int64          `gorm:"column:end_ts;type:bigint;comment:结束时间;not null;default:0;" json:"end_ts"`                           // 结束时间
	JobRole        string         `gorm:"column:job_role;type:varchar(256);comment:担任职务;not null;" json:"job_role"`                           // 担任职务
	Description    string         `gorm:"column:description;type:text;comment:经历描述;not null;" json:"description"`                             // 经历描述
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (CareerExperience) TableName() string {
	return "career_experience"
}

func CreateCareerExperience(ctx context.Context, db *gorm.DB, careerExperience *CareerExperience) error {
	if err := db.WithContext(ctx).Create(careerExperience).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] create career experience failed, %v", err)
		return err
	}
	return nil
}

func FindCareerExperienceById(ctx context.Context, db *gorm.DB, id int64) (*CareerExperience, error) {
	var careerExperience CareerExperience
	err := db.WithContext(ctx).Model(&CareerExperience{}).
		Where("id = ?", id).
		First(&careerExperience).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find career experience by id failed, %v", err)
		return nil, err
	}
	return &careerExperience, nil
}

func FindCareerExperienceByIdAndUserId(ctx context.Context, db *gorm.DB, id, userId int64) (*CareerExperience, error) {
	var careerExperience CareerExperience
	err := db.WithContext(ctx).Model(&CareerExperience{}).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		First(&careerExperience).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find career experience failed, %v", err)
		return nil, err
	}
	return &careerExperience, nil
}

func UpdateCareerExperience(ctx context.Context, db *gorm.DB, id int64, careerExperience *CareerExperience) error {
	err := db.WithContext(ctx).Model(careerExperience).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"begin_ts":        careerExperience.StartTS,
			"end_ts":          careerExperience.EndTS,
			"job_role":        careerExperience.JobRole,
			"description":     careerExperience.Description,
			"experience_name": careerExperience.ExperienceName,
		}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update career experience failed, %v", err)
		return err
	}
	return nil
}

func FindCareerExperienceByUserId(ctx context.Context, db *gorm.DB, userId int64) ([]*CareerExperience, error) {
	var careerExperiences []*CareerExperience
	err := db.WithContext(ctx).Model(&CareerExperience{}).
		Where("user_id = ?", userId).
		Find(&careerExperiences).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find career experience by user id failed, %v", err)
		return nil, err
	}
	return careerExperiences, nil
}

func DeleteCareerExperienceById(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Model(&CareerExperience{}).
		Where("id = ?", id).
		Delete(&CareerExperience{}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] delete career experience by id failed, %v", err)
		return err
	}
	return nil
}
