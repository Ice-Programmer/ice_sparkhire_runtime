package recruitment

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/career"
	"ice_sparkhire_runtime/service/company"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/utils"
)

func QueryRecruitmentPage(ctx context.Context, req *sparkruntime.QueryRecruitmentPageRequest) (*sparkruntime.QueryRecruitmentPageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	recruitmentList, total, err := db.QueryRecruitmentPage(ctx, db.DB, pageSize, pageNum, req.GetCondition())
	if err != nil {
		return nil, err
	}

	recruitmentPageInfos, err := buildEvaluatePageInfo(ctx, recruitmentList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.QueryRecruitmentPageResponse{
		RecruitmentList: recruitmentPageInfos,
		Total:           total,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}

func buildEvaluatePageInfo(ctx context.Context, recruitmentList []*db.Recruitment) ([]*sparkruntime.RecruitmentInfo, error) {
	// company
	companyIdList := utils.MapStructListDistinct(recruitmentList, func(recruitment *db.Recruitment) int64 {
		return recruitment.CompanyId
	})

	companyInfoMap, err := company.BuildCompanyInfoMapByIds(ctx, companyIdList)
	if err != nil {
		return nil, err
	}

	// career
	careerIdList := utils.MapStructListDistinct(recruitmentList, func(recruitment *db.Recruitment) int64 {
		return recruitment.CareerId
	})

	careerInfoMap, err := career.BuildCareerMapByIds(ctx, careerIdList)
	if err != nil {
		return nil, err
	}

	recruitmentInfoList := make([]*sparkruntime.RecruitmentInfo, 0, len(recruitmentList))
	for _, recruitment := range recruitmentList {
		geoInfo, err := geo.BuildGeoDetailInfo(ctx,
			recruitment.FirstGeoLevelId,
			recruitment.SecondGeoLevelId,
			recruitment.ThirdGeoLevelId,
			recruitment.ForthGeoLevelId,
			recruitment.Address,
			recruitment.Longitude,
			recruitment.Latitude,
		)
		if err != nil {
			klog.CtxErrorf(ctx, "build recruitment info failed, %v", err.Error())
			continue
		}

		tagList, err := db.FindTagsByObjIdAndObjType(ctx, db.DB, recruitment.ID, sparkruntime.TagObjType_Recruitment)
		if err != nil {
			klog.CtxErrorf(ctx, "find recruitment info failed, %v", err.Error())
			continue
		}

		recruitmentInfoList = append(recruitmentInfoList, &sparkruntime.RecruitmentInfo{
			Id:              recruitment.ID,
			Name:            recruitment.Name,
			CompanyInfo:     companyInfoMap[recruitment.CompanyId],
			CareerInfo:      careerInfoMap[recruitment.CareerId],
			Description:     recruitment.Description,
			Requirement:     recruitment.Requirement,
			EducationStatus: sparkruntime.EducationStatus(recruitment.EducationType),
			JobType:         sparkruntime.JobType(recruitment.JobType),
			GeoInfo:         geoInfo,
			SalaryInfo: &sparkruntime.SalaryInfo{
				SalaryUpper:   recruitment.SalaryUpper,
				SalaryLower:   recruitment.SalaryLower,
				CurrencyType:  sparkruntime.SalaryCurrencyType(recruitment.CurrencyType),
				FrequencyType: sparkruntime.SalaryFrequencyType(recruitment.FrequencyType),
			},
			TagInfo: utils.MapStructList(tagList, func(tag *db.Tag) *sparkruntime.TagInfo {
				return &sparkruntime.TagInfo{
					Id:      tag.Id,
					TagName: tag.TagName,
				}
			}),
		})
	}

	return recruitmentInfoList, nil
}
