package db

import (
	"context"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type ChatSession struct {
	Id              int64          `gorm:"column:id;type:bigint;comment:会话id;primaryKey;" json:"id"`                                     // 会话id
	RecruitmentId   int64          `gorm:"column:recruitment_id;type:bigint;comment:岗位id;not null;" json:"recruitment_id"`               // 岗位id
	CompanyId       int64          `gorm:"column:company_id;type:bigint;comment:公司id;not null;" json:"company_id"`                       // 公司id
	CandidateUserId int64          `gorm:"column:candidate_user_id;type:bigint;comment:求职者id;not null;" json:"candidate_user_id"`        // 求职者id
	HrUserId        int64          `gorm:"column:hr_user_id;type:bigint;comment:hr用户id;not null;" json:"hr_user_id"`                     // hr用户id
	LastMessageId   int64          `gorm:"column:last_message_id;type:bigint;comment:最后一条消息id;default:NULL;" json:"last_message_id"`     // 最后一条消息id
	LastMessage     string         `gorm:"column:last_message;type:varchar(500);comment:最后消息内容;default:NULL;" json:"last_message"`       // 最后消息内容
	LastMessageType int8           `gorm:"column:last_message_type;type:tinyint;comment:最后消息类型;default:1;" json:"last_message_type"`     // 最后消息类型
	LastMessageTime time.Time      `gorm:"column:last_message_time;type:datetime;comment:最后消息时间;default:NULL;" json:"last_message_time"` // 最后消息时间
	CandidateUnread int32          `gorm:"column:candidate_unread;type:int;comment:求职者未读数;default:0;" json:"candidate_unread"`           // 求职者未读数
	HrUnread        int32          `gorm:"column:hr_unread;type:int;comment:hr未读数;default:0;" json:"hr_unread"`                          // hr未读数
	Status          int8           `gorm:"column:status;type:tinyint;comment:状态 1正常 2已结束 3已屏蔽;default:1;" json:"status"`                 // 状态 1正常 2已结束 3已屏蔽
	CreateTime      time.Time      `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"create_time"`
	UpdateTime      time.Time      `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"update_time"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"` // 删除时间
}

func (*ChatSession) TableName() string {
	return "chat_session"
}

func ListChatSessionByCandidateId(ctx context.Context, db *gorm.DB, candidateId int64) ([]*ChatSession, error) {
	var sessionList []*ChatSession
	err := db.WithContext(ctx).Model(&ChatSession{}).
		Where("candidate_user_id = ?", candidateId).
		Find(&sessionList).Error
	if err != nil {
		return nil, err
	}
	return sessionList, nil
}

func FindChatSessionById(ctx context.Context, db *gorm.DB, id int64) (*ChatSession, error) {
	var session ChatSession
	err := db.WithContext(ctx).
		Model(&ChatSession{}).
		Where("id = ?", id).
		First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func EditChatSessionLastMessage(ctx context.Context, db *gorm.DB, sessionId int64, content string) error {
	err := db.WithContext(ctx).Model(&ChatSession{}).
		Where("id = ?", sessionId).
		Update("last_message", content).
		Update("last_message_time", time.Now()).
		Update("last_message_type", sparkhire_runtime.MessageType_Text).Error
	if err != nil {
		return err
	}
	return nil
}
