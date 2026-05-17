package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type RecruitmentApplication struct {
	Id            int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                                                                                      // id
	UserId        int64          `gorm:"column:user_id;type:bigint;comment:用户 id;not null;" json:"user_id"`                                                                                           // 用户 id
	RecruitmentId int64          `gorm:"column:recruitment_id;type:bigint;comment:招聘岗位 id;not null;" json:"recruitment_id"`                                                                           // 招聘岗位 id
	CompanyId     int64          `gorm:"column:company_id;type:bigint;comment:公司 id (对应 company 表 id);not null;" json:"company_id"`                                                                   // 公司 id (对应 company 表 id)
	Status        int8           `gorm:"column:status;type:tinyint;comment:当前推进状态：1-已投递/新投递, 2-简历筛选中, 3-约面/面试中, 4-待发Offer, 5-已发Offer, 6-已录用/入职, 7-不合适/淘汰, 8-求职者放弃;not null;default:1;" json:"status"` // 当前推进状态：1-已投递/新投递, 2-简历筛选中, 3-约面/面试中, 4-待发Offer, 5-已发Offer, 6-已录用/入职, 7-不合适/淘汰, 8-求职者放弃
	Remark        string         `gorm:"column:remark;type:varchar(512);comment:最新进度备注（如淘汰原因、面试评语摘要）;" json:"remark"`                                                                                 // 最新进度备注（如淘汰原因、面试评语摘要）
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;comment:投递时间/创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`                                                     // 投递时间/创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`                                                          // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                                                                             // 删除时间
}

func (RecruitmentApplication) TableName() string {
	return "recruitment_application"
}

func ExistRecruitmentCandidateApplication(ctx context.Context, db *gorm.DB, userId, recruitmentId int64) (bool, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&RecruitmentApplication{}).
		Where("user_id = ? AND recruitment_id = ?", userId, recruitmentId).
		Count(&count).Error

	if err != nil {
		klog.CtxErrorf(ctx, "[RecruitmentApplicationDB] exist recruitment application error: %v", err)
		return false, err
	}

	return count > 0, nil
}

func CreateRecruitmentApplication(ctx context.Context, db *gorm.DB, application *RecruitmentApplication) error {
	if err := db.WithContext(ctx).Model(RecruitmentApplication{}).Create(application).Error; err != nil {
		klog.CtxErrorf(ctx, "[RecruitmentApplicationDB] create recruitment application error: %v", err)
		return err
	}
	return nil
}
