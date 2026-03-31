package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type InterviewSchedule struct {
	Id                int64                        `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                                           // id
	CandidateId       int64                        `gorm:"column:candidate_id;type:bigint;comment:用户 id;not null;" json:"candidate_id"`                                      // 用户 id
	CreatorId         int64                        `gorm:"column:creator_id;type:bigint;comment:创建 id;not null;" json:"creator_id"`                                          // 创建 id
	RecruitmentId     int64                        `gorm:"column:recruitment_id;type:bigint;comment:招聘 id;not null;" json:"recruitment_id"`                                  // 招聘 id
	CompanyId         int64                        `gorm:"column:company_id;type:bigint;comment:公司 id;not null;" json:"company_id"`                                          // 公司 id
	InterviewTs       int64                        `gorm:"column:interview_ts;type:bigint;comment:面试开始时间;not null;" json:"interview_ts"`                                     // 面试开始时间
	InterviewDate     string                       `gorm:"column:interview_date;type:varchar(64);comment:面试时间 yyyy-MM-dd;not null;" json:"interview_date"`                   // 面试时间 yyyy-MM-dd
	InterviewDuration int32                        `gorm:"column:interview_duration;type:int;comment:面试时间（分钟）;not null;" json:"interview_duration"`                          // 面试时间（分钟）
	InterviewType     sparkruntime.InterviewType   `gorm:"column:interview_type;type:tinyint;comment:面试类型：1-视频面试, 2-线下面试, 3-电话面试;not null;default:1;" json:"interview_type"` // 面试类型：1-视频面试, 2-线下面试, 3-电话面试
	InterviewLink     string                       `gorm:"column:interview_link;type:varchar(512);comment:面试会议链接;" json:"interview_link"`                                    // 面试会议链接
	Status            sparkruntime.InterviewStatus `gorm:"column:status;type:tinyint;comment:状态：1-未开始, 2-进行中, 3-已结束/历史, 4-已取消;not null;default:1;" json:"status"`            // 状态：1-未开始, 2-进行中, 3-已结束/历史, 4-已取消
	CreatedAt         time.Time                    `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`               // 创建时间
	UpdatedAt         time.Time                    `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`               // 更新时间
	DeletedAt         gorm.DeletedAt               `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                                  // 删除时间
}

func (InterviewSchedule) TableName() string {
	return "interview_schedule"
}

func ListInterviewScheduleByCandidateId(ctx context.Context, db *gorm.DB, candidateId int64) ([]*InterviewSchedule, error) {
	var interviewList []*InterviewSchedule
	err := db.WithContext(ctx).Model(&InterviewSchedule{}).
		Where("candidate_id = ?", candidateId).
		Find(&interviewList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[InterviewSchedule] ListInterviewScheduleByCandidateId error: %v", err)
		return nil, err
	}

	return interviewList, nil
}
