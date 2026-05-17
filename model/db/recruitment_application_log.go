package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type RecruitmentApplicationLog struct {
	Id            int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                                      // id
	ApplicationId int64          `gorm:"column:application_id;type:bigint;comment:职位申请 id (对应 job_application 表 id);not null;" json:"application_id"` // 职位申请 id (对应 job_application 表 id)
	FromStatus    int8           `gorm:"column:from_status;type:tinyint;comment:变更前状态(0表示初始投递);not null;default:0;" json:"from_status"`               // 变更前状态(0表示初始投递)
	ToStatus      int8           `gorm:"column:to_status;type:tinyint;comment:变更后状态;not null;" json:"to_status"`                                      // 变更后状态
	OperatorId    int64          `gorm:"column:operator_id;type:bigint;comment:操作人 id (对应 user 表 id，可以是 HR 或系统或求职者自己);not null;" json:"operator_id"`  // 操作人 id (对应 user 表 id，可以是 HR 或系统或求职者自己)
	Remark        string         `gorm:"column:remark;type:varchar(1024);comment:阶段操作备注/评语/原因;" json:"remark"`                                        // 阶段操作备注/评语/原因
	CreatedAt     gorm.DeletedAt `gorm:"column:created_at;type:datetime;comment:操作时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`          // 操作时间
}

func (RecruitmentApplicationLog) TableName() string {
	return "recruitment_application_log"
}

func CreateRecruitmentApplicationLog(ctx context.Context, db *gorm.DB, log *RecruitmentApplicationLog) error {
	if err := db.WithContext(ctx).Model(&RecruitmentApplicationLog{}).Create(log).Error; err != nil {
		return fmt.Errorf("create recruitment application %w", err)
	}
	return nil
}
