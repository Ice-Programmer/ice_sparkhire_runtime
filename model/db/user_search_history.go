package db

import (
	"context"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type UserSearchHistory struct {
	Id            int64                        `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	UserId        int64                        `gorm:"column:user_id;type:bigint;comment:创建用户 id;not null;" json:"user_id"`                                // 创建用户 id
	SearchContent string                       `gorm:"column:search_content;type:varchar(256);comment:搜索内容;not null;" json:"search_content"`               // 搜索内容
	Type          sparkruntime.UserHistoryType `gorm:"column:type;type:tinyint;comment:搜索类型 1-recruitment 2-post;not null;default:1;" json:"type"`         // 搜索类型 1-recruitment 2-post
	CreatedAt     time.Time                    `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	DeletedAt     gorm.DeletedAt               `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (*UserSearchHistory) TableName() string {
	return "user_search_history"
}

func CreateUserSearchHistory(ctx context.Context, db *gorm.DB, searchHistory *UserSearchHistory) error {
	if err := db.WithContext(ctx).Model(&UserSearchHistory{}).Create(searchHistory).Error; err != nil {
		return err
	}

	return nil
}

func ListUserSearchHistoryByUserId(ctx context.Context, db *gorm.DB, userType sparkruntime.UserHistoryType, userId int64, lastNum int) ([]*UserSearchHistory, error) {
	var historyList []*UserSearchHistory
	err := db.WithContext(ctx).Model(&UserSearchHistory{}).
		Where("user_id = ?", userId).
		Where("type = ?", userType).
		Order("created_at DESC").
		Limit(lastNum).
		Find(&historyList).Error
	if err != nil {
		return nil, err
	}

	return historyList, nil
}
