package db

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type UserFavorite struct {
	Id         int64          `gorm:"column:id;type:bigint;comment:主键 id;primaryKey;" json:"id"`                                          // 主键 id
	UserId     int64          `gorm:"column:user_id;type:bigint;comment:用户 id;not null;" json:"user_id"`                                  // 用户 id
	TargetType int8           `gorm:"column:target_type;type:tinyint;comment:收藏目标类型: 1-公司, 2-职位, 3-文章等;not null;" json:"target_type"`     // 收藏目标类型: 1-公司, 2-职位, 3-文章等
	TargetId   int64          `gorm:"column:target_id;type:bigint;comment:目标 id (公司id/职位id等);not null;" json:"target_id"`                 // 目标 id (公司id/职位id等)
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (UserFavorite) TableName() string {
	return "user_favorite"
}

func CreateUserFavor(ctx context.Context, db *gorm.DB, userFavor *UserFavorite) error {
	err := db.WithContext(ctx).Model(&UserFavorite{}).
		Create(userFavor).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] create userFavor error: %v", err)
		return err
	}
	return nil
}

func DeleteUserFavor(ctx context.Context, db *gorm.DB, userId, targetId int64, targetType sparkruntime.TargetType) error {
	result := db.WithContext(ctx).Model(&UserFavorite{}).
		Where("target_id = ?", targetId).
		Where("target_type = ?", targetType).
		Where("user_id = ?", userId).
		Delete(&UserFavorite{})
	if err := result.Error; err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] delete userFavorite error: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("favorite record is not exist")
	}

	return nil
}

func HasFavor(ctx context.Context, db *gorm.DB, userId int64, targetId int64, targetType sparkruntime.TargetType) bool {
	var exists int
	err := db.WithContext(ctx).Model(&UserFavorite{}).
		Select("1").
		Where("user_id = ?", userId).
		Where("target_id = ?", targetId).
		Where("target_type = ?", targetType).
		Limit(1).Scan(&exists).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] has favorite err: %v", err)
		return false
	}

	return exists == 1
}
