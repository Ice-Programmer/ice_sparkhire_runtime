package recruitment

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/company"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/utils"
)

func FetchRecruitmentInfo(ctx context.Context, req *sparkruntime.FetchRecruitmentInfoRequest) (*sparkruntime.FetchRecruitmentInfoResponse, error) {
	// 1. validate id
	if req.GetRecruitmentId() <= 0 {
		return nil, fmt.Errorf("invalid recruitment id")
	}

	// 2. find recruitment
	recruitment, err := db.FindRecruitmentById(ctx, db.DB, req.GetRecruitmentId())
	if err != nil {
		return nil, err
	}

	// 3. find company
	companyDB, err := db.FindCompanyById(ctx, db.DB, recruitment.CompanyId)
	if err != nil {
		return nil, err
	}

	companyInfo, err := company.BuildCompanyInfo(ctx, companyDB)
	if err != nil {
		return nil, err
	}

	// 4. career
	careerDB, err := db.FindCareerById(ctx, db.DB, recruitment.CareerId)
	if err != nil {
		return nil, err
	}

	// 5. geo info
	geoDetailInfo, err := geo.BuildGeoDetailInfo(ctx,
		recruitment.FirstGeoLevelId,
		recruitment.SecondGeoLevelId,
		recruitment.ThirdGeoLevelId,
		recruitment.ForthGeoLevelId,
		recruitment.Address,
		recruitment.Longitude,
		recruitment.Latitude,
	)
	if err != nil {
		return nil, err
	}

	tagList, err := db.FindTagsByObjIdAndObjType(ctx, db.DB, recruitment.ID, sparkruntime.TagObjType_Recruitment)
	if err != nil {
		klog.CtxErrorf(ctx, "find recruitment info failed, %v", err.Error())
		return nil, err
	}

	recruitmentInfo := &sparkruntime.RecruitmentInfo{
		Id:          recruitment.ID,
		Name:        recruitment.Name,
		CompanyInfo: companyInfo,
		CareerInfo: &sparkruntime.CareerInfo{
			Id:           careerDB.Id,
			CareerName:   careerDB.CareerName,
			CareerTypeId: utils.Int64Ptr(careerDB.CareerType),
			CareerIcon:   careerDB.CareerIcon,
			Description:  careerDB.Description,
		},
		Description:     recruitment.Description,
		Requirement:     recruitment.Requirement,
		EducationStatus: sparkruntime.EducationStatus(recruitment.EducationType),
		JobType:         sparkruntime.JobType(recruitment.JobType),
		GeoInfo:         geoDetailInfo,
		SalaryInfo: &sparkruntime.SalaryInfo{
			SalaryUpper:   recruitment.SalaryUpper,
			SalaryLower:   recruitment.SalaryLower,
			CurrencyType:  sparkruntime.SalaryCurrencyType(recruitment.CurrencyType),
			FrequencyType: sparkruntime.SalaryFrequencyType(recruitment.FrequencyType),
		},
		TagInfoList: utils.MapStructList(tagList, func(tag *db.Tag) *sparkruntime.TagInfo {
			return &sparkruntime.TagInfo{
				Id:      tag.Id,
				TagName: tag.TagName,
			}
		}),
	}

	return &sparkruntime.FetchRecruitmentInfoResponse{
		RecruitmentInfo: recruitmentInfo,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}
