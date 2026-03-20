package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func FindUserFavorByTargetIdAndUserId(ctx context.Context, db *gorm.DB, userId, targetId int64, targetType sparkruntime.TargetType) (*UserFavorite, error) {
	var userFavorite UserFavorite
	err := db.WithContext(ctx).Model(&UserFavorite{}).
		Where("user_id = ?", userId).
		Where("target_id = ?", targetId).
		Where("target_type = ?", targetType).
		Find(&userFavorite).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] find user favorite err: %v", err)
		return nil, err
	}

	return &userFavorite, nil
}

func UpsertUserFavor(ctx context.Context, db *gorm.DB, userFavor *UserFavorite) error {
	err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"},
			{Name: "target_type"},
			{Name: "target_id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"deleted_at": nil,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}),
	}).Create(userFavor).Error

	if err != nil {
		klog.CtxErrorf(ctx, "[UserFavorDB] create userFavor error: %v", err)
		return err
	}
	return nil
}
