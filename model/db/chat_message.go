package db

import (
	"context"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type ChatMessage struct {
	Id           int64                    `gorm:"column:id;type:bigint;comment:消息id;primaryKey;" json:"id"`                                                           // 消息id
	SessionId    int64                    `gorm:"column:session_id;type:bigint;comment:会话id;not null;" json:"session_id"`                                             // 会话id
	SenderId     int64                    `gorm:"column:sender_id;type:bigint;comment:发送者id;not null;" json:"sender_id"`                                              // 发送者id
	ReceiverId   int64                    `gorm:"column:receiver_id;type:bigint;comment:接收者id;not null;" json:"receiver_id"`                                          // 接收者id
	SenderType   sparkruntime.SenderType  `gorm:"column:sender_type;type:tinyint;comment:发送者类型 1候选人 2hr;not null;" json:"sender_type"`                                // 发送者类型 1候选人 2hr
	MessageType  sparkruntime.MessageType `gorm:"column:message_type;type:tinyint;comment:1文本 2图片 3文件 4简历 5岗位卡片 6面试邀请 7系统消息;not null;default:1;" json:"message_type"` // 1文本 2图片 3文件 4简历 5岗位卡片 6面试邀请 7系统消息
	Content      string                   `gorm:"column:content;type:text;comment:消息内容;" json:"content"`                                                              // 消息内容
	IsRead       int8                     `gorm:"column:is_read;type:tinyint;comment:是否已读;not null;default:0;" json:"is_read"`                                        // 是否已读
	RevokeStatus int8                     `gorm:"column:revoke_status;type:tinyint;comment:撤回状态;not null;default:0;" json:"revoke_status"`                            // 撤回状态
	SendStatus   int8                     `gorm:"column:send_status;type:tinyint;comment:1发送中 2发送成功 3发送失败;not null;default:1;" json:"send_status"`                    // 1发送中 2发送成功 3发送失败
	CreateTime   time.Time                `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"create_time"`
	DeletedAt    gorm.DeletedAt           `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"` // 删除时间
}

func (*ChatMessage) TableName() string {
	return "chat_message"
}

func QueryChatMessagePage(ctx context.Context, db *gorm.DB, pageSize, pageNum int32, sessionId int64) ([]*ChatMessage, int64, error) {
	query := db.WithContext(ctx).Model(&ChatMessage{}).Where("session_id = ?", sessionId)

	var total int64
	if err := query.WithContext(ctx).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := int((pageNum - 1) * pageSize)
	var messageList []*ChatMessage
	err := query.WithContext(ctx).
		Offset(offset).
		Limit(int(pageSize)).
		Order("create_time DESC").
		Find(&messageList).Error
	if err != nil {
		return nil, 0, err
	}

	return messageList, total, nil
}

func CreateChatMessage(ctx context.Context, db *gorm.DB, chatMessage *ChatMessage) error {
	if err := db.WithContext(ctx).Model(&ChatMessage{}).Create(chatMessage).Error; err != nil {
		return err
	}
	return nil
}
