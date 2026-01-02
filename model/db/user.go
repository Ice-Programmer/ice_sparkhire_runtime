package db

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

type User struct {
	Id         int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`
	Username   string         `gorm:"column:username;type:varchar(128);comment:用户昵称;not null;" json:"username"`
	UserAvatar string         `gorm:"column:user_avatar;type:varchar(128);comment:用户头像;not null;" json:"user_avatar"`
	Email      string         `gorm:"column:email;type:varchar(256);comment:邮箱;not null;" json:"email"`
	Gender     int8           `gorm:"column:gender;type:tinyint;comment:0-女 1-男;not null;default:1;" json:"gender"`
	UserRole   int8           `gorm:"column:user_role;type:tinyint;comment:用户角色（1-visitor 2-candidate 3-HR 4-admin）;not null;default:0;" json:"user_role"`
	Status     int8           `gorm:"column:status;type:tinyint;comment:用户状态(0-正常 1-封禁);not null;default:0;" json:"status"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;" json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find user by email %s error: %v", email, err)
		return nil, err
	}

	return &user, nil
}

func FindUserById(ctx context.Context, db *gorm.DB, id int64) (*User, error) {
	var user User
	err := db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find user by id %v error: %v", id, err)
		return nil, err
	}
	return &user, err
}

func UpdateUserRole(ctx context.Context, db *gorm.DB, role sparkruntime.UserRole, id int64) error {
	err := db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).
		Update("user_role", role).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update user role %v error: %v", role, err)
		return err
	}
	return nil
}

func SaveUser(ctx context.Context, db *gorm.DB, user *User) error {
	err := db.WithContext(ctx).Save(user).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] save user error: %v", err)
		return err
	}
	return nil
}

func UpdateUserById(ctx context.Context, db *gorm.DB, id int64, updateMap map[string]interface{}) error {
	err := db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).
		Updates(updateMap).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update user id %v error: %v", id, err)
		return err
	}
	return nil
}
