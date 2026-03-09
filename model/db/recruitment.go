package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

const RecruitmentTableName = "recruitment"

type Recruitment struct {
	ID               int64          `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	Name             string         `gorm:"column:name;type:varchar(512);comment:岗位招聘标题;not null;" json:"name"`                                 // 岗位招聘标题
	UserId           int64          `gorm:"column:user_id;type:bigint;comment:岗位发布者id;not null;" json:"user_id"`                                // 岗位发布者id
	CompanyId        int64          `gorm:"column:company_id;type:bigint;comment:公司id;not null;" json:"company_id"`                             // 公司id
	CareerId         int64          `gorm:"column:career_id;type:bigint;comment:职业id;not null;" json:"career_id"`                               // 职业id
	Description      string         `gorm:"column:description;type:text;comment:职位详情;not null;" json:"description"`                             // 职位详情
	Requirement      string         `gorm:"column:requirement;type:text;comment:职位要求;not null;" json:"requirement"`                             // 职位要求
	EducationType    int8           `gorm:"column:education_type;type:tinyint;comment:最低学历要求;" json:"education_type"`                           // 最低学历要求
	JobType          int8           `gorm:"column:job_type;type:tinyint;comment:职业类型（实习、兼职、春招等）;not null;" json:"job_type"`                     // 职业类型（实习、兼职、春招等）
	ApplyCount       int32          `gorm:"column:apply_count;type:int;comment:投递次数;not null;default:0;" json:"apply_count"`                    // 投递次数
	FavoriteCount    int32          `gorm:"column:favorite_count;type:int;comment:收藏次数;not null;default:0;" json:"favorite_count"`              // 收藏次数
	FirstGeoLevelId  int64          `gorm:"column:first_geo_level_id;type:bigint;comment:一级地理位置 id;not null;" json:"first_geo_level_id"`        // 一级地理位置 id
	SecondGeoLevelId int64          `gorm:"column:second_geo_level_id;type:bigint;comment:二级地理位置 id;not null;" json:"second_geo_level_id"`      // 二级地理位置 id
	ThirdGeoLevelId  int64          `gorm:"column:third_geo_level_id;type:bigint;comment:三级地理位置 id;not null;" json:"third_geo_level_id"`        // 三级地理位置 id
	ForthGeoLevelId  int64          `gorm:"column:forth_geo_level_id;type:bigint;comment:四级地理位置 id;not null;" json:"forth_geo_level_id"`        // 四级地理位置 id
	Address          string         `gorm:"column:address;type:varchar(512);comment:具体地址;not null;" json:"address"`                             // 具体地址
	Latitude         float64        `gorm:"column:latitude;type:decimal(10, 7);comment:纬度;" json:"latitude"`                                    // 纬度
	Longitude        float64        `gorm:"column:longitude;type:decimal(10, 7);comment:经度;" json:"longitude"`                                  // 经度
	SalaryUpper      int32          `gorm:"column:salary_upper;type:int;comment:薪水上限;not null;default:0;" json:"salary_upper"`                  // 薪水上限
	SalaryLower      int32          `gorm:"column:salary_lower;type:int;comment:薪水下限;not null;default:0;" json:"salary_lower"`                  // 薪水下限
	CurrencyType     int32          `gorm:"column:currency_type;type:int;comment:薪水货币类型;not null;default:1;" json:"currency_type"`              // 薪水货币类型
	FrequencyType    int8           `gorm:"column:frequency_type;type:tinyint;comment:类型;not null;default:1;" json:"frequency_type"`            // 类型
	Status           int8           `gorm:"column:status;type:tinyint;comment:招聘状态（0 - 招聘中 1 - 结束招聘）;not null;default:0;" json:"status"`        // 招聘状态（0 - 招聘中 1 - 结束招聘）
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"`                                    // 删除时间
}

func (Recruitment) TableName() string {
	return RecruitmentTableName
}

func CreateRecruitment(ctx context.Context, db *gorm.DB, recruitment *Recruitment) error {
	if err := db.WithContext(ctx).Create(recruitment).Error; err != nil {
		klog.CtxErrorf(ctx, "[db] create recruitment err: %v", err)
		return err
	}

	return nil
}

func FindRecruitmentById(ctx context.Context, db *gorm.DB, id int64) (*Recruitment, error) {
	var recruitment Recruitment

	err := db.WithContext(ctx).Model(&Recruitment{}).
		Where("id = ?", id).
		First(&recruitment).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] find recruitment err: %v", err)
		return nil, err
	}

	return &recruitment, nil
}

func QueryRecruitmentPage(ctx context.Context, db *gorm.DB, pageSize, pageNum int32, condition *sparkruntime.RecruitmentCondition) ([]*Recruitment, int64, error) {
	query := buildRecruitmentCondition(db, condition)

	var total int64
	if err := query.WithContext(ctx).Count(&total).Error; err != nil {
		klog.CtxErrorf(ctx, "[db] count recruitment err: %v", err)
		return nil, 0, err
	}

	offset := int((pageNum - 1) * pageSize)
	var recruitmentList []*Recruitment
	err := query.WithContext(ctx).
		Offset(offset).
		Limit(int(pageSize)).
		Order("created_at DESC").
		Find(&recruitmentList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[db] find recruitment err: %v", err)
		return nil, 0, err
	}

	return recruitmentList, total, nil
}

func buildRecruitmentCondition(db *gorm.DB, condition *sparkruntime.RecruitmentCondition) *gorm.DB {
	db = db.Model(&Recruitment{})

	if condition == nil {
		return db
	}

	if condition.IsSetSalaryUpper() {
		db = db.Where("salary_upper <= ?", condition.SalaryUpper)
	}

	if condition.IsSetSalaryLower() {
		db = db.Where("salary_lower >= ?", condition.SalaryLower)
	}

	if condition.IsSetJobType() {
		db = db.Where("job_type = ?", condition.JobType)
	}

	if condition.IsSetEducationStatus() {
		db = db.Where("education_type = ?", condition.EducationStatus)
	}

	if condition.IsSetCompanyId() {
		db = db.Where("company_id = ?", condition.CompanyId)
	}

	if condition.IsSetCareerId() {
		db = db.Where("career_id = ?", condition.CareerId)
	}

	return db
}
