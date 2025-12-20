package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type GeoFirstLevel struct {
	Id        int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	GeoName   string    `gorm:"column:geo_name;type:varchar(128);comment:地理名称;not null;" json:"geo_name"`                           // 地理名称
	Code      string    `gorm:"column:code;type:varchar(128);comment:code;not null;" json:"code"`                                   // code
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (GeoFirstLevel) TableName() string {
	return "geo_first_level"
}

type GeoSecondLevel struct {
	Id              int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	GeoName         string    `gorm:"column:geo_name;type:varchar(128);comment:地理名称;not null;" json:"geo_name"`                           // 地理名称
	Code            string    `gorm:"column:code;type:varchar(128);comment:code;not null;" json:"code"`                                   // code
	FirstGeoLevelId int64     `gorm:"column:first_geo_level_id;type:bigint;comment:一级地理位置 id;not null;" json:"first_geo_level_id"`        // 一级地理位置 id
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (GeoSecondLevel) TableName() string {
	return "geo_second_level"
}

type GeoThirdLevel struct {
	Id               int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	GeoName          string    `gorm:"column:geo_name;type:varchar(128);comment:地理名称;not null;" json:"geo_name"`                           // 地理名称
	Code             string    `gorm:"column:code;type:varchar(128);comment:code;not null;" json:"code"`                                   // code
	SecondGeoLevelId int64     `gorm:"column:second_geo_level_id;type:bigint;comment:二级地理位置 id;not null;" json:"second_geo_level_id"`      // 二级地理位置 id
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (GeoThirdLevel) TableName() string {
	return "geo_third_level"
}

type GeoForthLevel struct {
	Id              int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                    // id
	GeoName         string    `gorm:"column:geo_name;type:varchar(128);comment:地理名称;not null;" json:"geo_name"`                           // 地理名称
	Code            string    `gorm:"column:code;type:varchar(128);comment:code;not null;" json:"code"`                                   // code
	ThirdGeoLevelId int64     `gorm:"column:third_geo_level_id;type:bigint;comment:三级地理位置 id;not null;" json:"third_geo_level_id"`        // 三级地理位置 id
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (GeoForthLevel) TableName() string {
	return "geo_forth_level"
}

func ListAllFirstGeo(ctx context.Context, db *gorm.DB) ([]*GeoFirstLevel, error) {
	var list []*GeoFirstLevel
	err := db.Model(&GeoFirstLevel{}).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] list all first geo level error: %v", err)
		return nil, err
	}
	return list, err
}

func ListSecondGeoByFirstGeoId(ctx context.Context, db *gorm.DB, firstGeoId int64) ([]*GeoSecondLevel, error) {
	var list []*GeoSecondLevel
	err := db.WithContext(ctx).Where("first_geo_level_id = ?", firstGeoId).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] list all second geo level error: %v", err)
		return nil, err
	}
	return list, err
}

func ListThirdGeoBySecondGeoId(ctx context.Context, db *gorm.DB, secondGeoId int64) ([]*GeoThirdLevel, error) {
	var list []*GeoThirdLevel
	err := db.WithContext(ctx).Where("second_geo_level_id = ?", secondGeoId).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] list all third geo level error: %v", err)
		return nil, err
	}
	return list, err
}

func ListForthGeoByThirdGeoId(ctx context.Context, db *gorm.DB, thirdGeoId int64) ([]*GeoForthLevel, error) {
	var list []*GeoForthLevel
	err := db.WithContext(ctx).
		Where("third_geo_level_id = ?", thirdGeoId).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] list all forth geo level error: %v", err)
		return nil, err
	}
	return list, err
}

func FindFirstGeoById(ctx context.Context, db *gorm.DB, id int64) (*GeoFirstLevel, error) {
	var geo *GeoFirstLevel
	err := db.WithContext(ctx).Where("id = ?", id).First(&geo).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] get first geo level error: %v", err)
		return nil, err
	}
	return geo, err
}

func FindSecondGeoById(ctx context.Context, db *gorm.DB, id int64) (*GeoSecondLevel, error) {
	var geo *GeoSecondLevel
	err := db.WithContext(ctx).Where("id = ?", id).First(&geo).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] get second geo level error: %v", err)
		return nil, err
	}
	return geo, err
}

func FindThirdLevelById(ctx context.Context, db *gorm.DB, id int64) (*GeoThirdLevel, error) {
	var geo *GeoThirdLevel
	err := db.WithContext(ctx).Where("id = ?", id).First(&geo).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] get third geo level error: %v", err)
		return nil, err
	}
	return geo, err
}

func FindForthLevelById(ctx context.Context, db *gorm.DB, id int64) (*GeoForthLevel, error) {
	var geo *GeoForthLevel
	err := db.WithContext(ctx).Where("id = ?", id).First(&geo).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] get forth geo level error: %v", err)
		return nil, err
	}
	return geo, err
}
