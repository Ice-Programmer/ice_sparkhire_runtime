package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const CompanyTableName = "company"

type Company struct {
	ID               int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	CompanyName      string         `gorm:"column:company_name;type:varchar(256);comment:公司名称;not null;" json:"company_name"`                   // 公司名称
	CreateUserId     int64          `gorm:"column:create_user_id;type:bigint;comment:创建用户 id;not null;" json:"create_user_id"`                  // 创建用户 id
	Description      string         `gorm:"column:description;type:text;comment:公司介绍;not null;" json:"description"`                             // 公司介绍
	FavoriteCount    int32          `gorm:"column:favorite_count;type:int;comment:收藏次数;not null;default:0;" json:"favorite_count"`              // 收藏次数
	Logo             string         `gorm:"column:logo;type:varchar(256);comment:公司 Logo;not null;" json:"logo"`                                // 公司 Logo
	BackgroundImg    string         `gorm:"column:background_img;type:varchar(256);comment:公司背景图片;" json:"background_img"`                      // 公司背景图片
	CompanyImgList   string         `gorm:"column:company_img_list;type:varchar(1024);comment:公司图片;default:[];" json:"company_img_list"`        // 公司图片
	IndustryId       int64          `gorm:"column:industry_id;type:bigint;comment:公司产业;not null;" json:"industry_id"`                           // 公司产业
	FirstGeoLevelId  int64          `gorm:"column:first_geo_level_id;type:bigint;comment:一级地理位置 id;not null;" json:"first_geo_level_id"`        // 一级地理位置 id
	SecondGeoLevelId int64          `gorm:"column:second_geo_level_id;type:bigint;comment:二级地理位置 id;not null;" json:"second_geo_level_id"`      // 二级地理位置 id
	ThirdGeoLevelId  int64          `gorm:"column:third_geo_level_id;type:bigint;comment:三级地理位置 id;not null;" json:"third_geo_level_id"`        // 三级地理位置 id
	ForthGeoLevelId  int64          `gorm:"column:forth_geo_level_id;type:bigint;comment:四级地理位置 id;not null;" json:"forth_geo_level_id"`        // 四级地理位置 id
	Address          string         `gorm:"column:address;type:varchar(512);comment:具体地址;not null;" json:"address"`                             // 具体地址
	Latitude         float64        `gorm:"column:latitude;type:decimal(10, 7);comment:纬度;" json:"latitude"`                                    // 纬度
	Longitude        float64        `gorm:"column:longitude;type:decimal(10, 7);comment:经度;" json:"longitude"`                                  // 经度
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (Company) TableName() string {
	return CompanyTableName
}

func CreateCompany(ctx context.Context, db *gorm.DB, company *Company) error {
	if err := db.WithContext(ctx).Create(company).Error; err != nil {
		klog.CtxErrorf(ctx, "[db] create company err: %v", err)
		return err
	}
	return nil
}

func FindCompanyByName(ctx context.Context, db *gorm.DB, name string) (*Company, error) {
	var company Company
	err := db.WithContext(ctx).Model(&Company{}).
		Where("company_name = ?", name).
		First(&company).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] find company err: %v", err)
		return nil, err
	}
	return &company, nil
}

func FindCompanyById(ctx context.Context, db *gorm.DB, id int64) (*Company, error) {
	var company Company
	err := db.WithContext(ctx).Model(&Company{}).
		Where("id = ?", id).
		First(&company).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] find company err: %v", err)
		return nil, err
	}
	return &company, nil
}

func UpdateCompany(ctx context.Context, db *gorm.DB, id int64, modifyCompany *Company) error {
	err := db.WithContext(ctx).Model(modifyCompany).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"company_name":        modifyCompany.CompanyName,
			"description":         modifyCompany.Description,
			"logo":                modifyCompany.Logo,
			"industry_id":         modifyCompany.IndustryId,
			"first_geo_level_id":  modifyCompany.FirstGeoLevelId,
			"second_geo_level_id": modifyCompany.SecondGeoLevelId,
			"third_geo_level_id":  modifyCompany.ThirdGeoLevelId,
			"forth_geo_level_id":  modifyCompany.ForthGeoLevelId,
			"address":             modifyCompany.Address,
			"latitude":            modifyCompany.Latitude,
			"longitude":           modifyCompany.Longitude,
		}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] update company err: %v", err)
		return err
	}
	return nil
}

func DeleteCompany(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&Company{}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] delete company err: %v", err)
		return err
	}
	return nil
}
