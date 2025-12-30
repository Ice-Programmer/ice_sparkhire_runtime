package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const (
	CandidateTableName = "candidate"
)

type Candidate struct {
	Id               int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;not null;" json:"id"`                                                // id
	UserId           int64          `gorm:"column:user_id;type:bigint;comment:用户id;not null;" json:"user_id"`                                               // 用户id
	Age              int32          `gorm:"column:age;type:int;comment:年龄;not null;default:20;" json:"age"`                                                 // 年龄
	Education        int32          `gorm:"column:education;type:int;comment:最高学历(1-本科 2-研究生 3-博士生 4-大专 5-高中 6-高中以下);not null;default:1;" json:"education"` // 最高学历(1-本科 2-研究生 3-博士生 4-大专 5-高中 6-高中以下)
	Phone            string         `gorm:"column:phone;type:varchar(128);comment:联系方式;not null;" json:"phone"`
	GraduationYear   int32          `gorm:"column:graduation_year;type:int;comment:毕业年份;not null;" json:"graduation_year"`                      // 毕业年份
	JobStatus        int8           `gorm:"column:job_status;type:tinyint;comment:求职状态;not null;" json:"job_status"`                            // 求职状态
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

func (Candidate) TableName() string {
	return CandidateTableName
}

func CreateCandidate(ctx context.Context, db *gorm.DB, candidate *Candidate) error {
	if err := db.WithContext(ctx).Save(candidate).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] save candidate error: %v", err)
		return err
	}

	return nil
}

func FindCandidateByUserId(ctx context.Context, db *gorm.DB, userId int64) (*Candidate, error) {
	var candidate Candidate
	err := db.WithContext(ctx).Where("user_id = ?", userId).First(&candidate).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find candidate by id %v error: %v", userId, err)
		return nil, err
	}

	return &candidate, nil
}

func FindCandidateById(ctx context.Context, db *gorm.DB, userId int64) (*Candidate, error) {
	var candidate Candidate
	err := db.WithContext(ctx).Where("id = ?", userId).First(&candidate).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find candidate by id %v error: %v", userId, err)
		return nil, err
	}
	return &candidate, nil
}

func UpdateCandidateByUserId(ctx context.Context, db *gorm.DB, userId int64, updateFields map[string]interface{}) error {
	err := db.WithContext(ctx).Model(&Candidate{}).
		Where("user_id = ?", userId).
		Updates(updateFields).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] update candidate error: %v", err)
		return err
	}

	return nil
}
