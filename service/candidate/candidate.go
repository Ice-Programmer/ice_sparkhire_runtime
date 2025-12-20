package candidate

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/utils"
)

func ValidateCandidate(candidate *sparkruntime.CandidateInfo) error {
	if candidate == nil {
		return fmt.Errorf("candidate is nil")
	}

	if candidate.GetAge() <= 0 || candidate.GetAge() >= 99 {
		return fmt.Errorf("invalid age")
	}

	if len(candidate.GetQualificationList()) > 20 {
		return fmt.Errorf("qualification list is too long")
	}

	if utils.NotContains(sparkruntime.JobStatusList, candidate.JobStatus) {
		return fmt.Errorf("job is invalide")
	}

	if utils.NotContains(sparkruntime.EducationStatusList, candidate.EducationStatus) {
		return fmt.Errorf("education is invalide")
	}

	if err := utils.ValidateGeoDetail(candidate.GetGeoDetail()); err != nil {
		return err
	}

	if err := utils.ValidateYear(candidate.GetGraduationYear()); err != nil {
		return err
	}

	return nil
}

func BuildCandidateDB(ctx context.Context, candidateInfo *sparkruntime.CandidateInfo, id int64) (*db.Candidate, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	return &db.Candidate{
		Id:               id,
		UserId:           userId,
		Age:              candidateInfo.GetAge(),
		Qualifications:   utils.MarshalString(candidateInfo.QualificationList),
		Education:        int32(candidateInfo.EducationStatus),
		GraduationYear:   candidateInfo.GraduationYear,
		JobStatus:        int8(candidateInfo.JobStatus),
		FirstGeoLevelId:  candidateInfo.GeoDetail.GetFirstGeoLevelId(),
		SecondGeoLevelId: candidateInfo.GeoDetail.GetSecondGeoLevelId(),
		ThirdGeoLevelId:  candidateInfo.GeoDetail.GetThirdGeoLevelId(),
		ForthGeoLevelId:  candidateInfo.GeoDetail.GetForthGeoLevelId(),
		Address:          candidateInfo.GeoDetail.GetAddress(),
		Latitude:         candidateInfo.GeoDetail.GetLatitude(),
		Longitude:        candidateInfo.GeoDetail.GetLongitude(),
	}, nil
}

func BuildCandidateInfo(ctx context.Context, candidate *db.Candidate) (*sparkruntime.CandidateInfo, error) {
	qualificationList := make([]string, 0)
	if err := sonic.Unmarshal([]byte(candidate.Qualifications), &qualificationList); err != nil {
		return nil, err
	}

	geoDetailInfo, err := geo.BuildGeoDetailInfo(ctx, candidate.FirstGeoLevelId, candidate.SecondGeoLevelId, candidate.ThirdGeoLevelId, candidate.ForthGeoLevelId, candidate.Address)
	if err != nil {
		return nil, err
	}

	tagList, err := db.FindTagsByObjIdAndObjType(ctx, db.DB, candidate.Id, int32(sparkruntime.TagObjType_Candidate))
	if err != nil {
		return nil, err
	}

	tagInfoList := utils.MapStructList(tagList, func(tag *db.Tag) *sparkruntime.TagInfo {
		return &sparkruntime.TagInfo{
			Id:      tag.Id,
			TagName: tag.TagName,
		}
	})

	return &sparkruntime.CandidateInfo{
		Age:               candidate.Age,
		QualificationList: qualificationList,
		JobStatus:         sparkruntime.JobStatus(candidate.JobStatus),
		GeoDetail:         geoDetailInfo,
		GraduationYear:    candidate.GraduationYear,
		EducationStatus:   sparkruntime.EducationStatus(candidate.Education),
		Id:                utils.Int64Ptr(candidate.Id),
		TagList:           tagInfoList,
	}, nil
}

func BindTags(ctx context.Context, objId int64, tagIdList []int64) (int, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return 0, err
	}

	// 1. candidate is existing
	candidate, err := db.FindCandidateById(ctx, db.DB, objId)
	if err != nil {
		return 0, err
	}
	if candidate.UserId != userId {
		return 0, fmt.Errorf("only can modify own tag")
	}

	tagList, err := db.FindTagByIds(ctx, db.DB, tagIdList)
	if err != nil {
		return 0, err
	}

	tagIdList = utils.MapStructList(tagList, func(tag *db.Tag) int64 {
		return tag.Id
	})

	tagRels, err := db.FindTagRelsByObjId(ctx, db.DB, objId, int32(sparkruntime.TagObjType_Candidate))
	if err != nil {
		return 0, err
	}

	tagRelIdList := utils.MapStructList(tagRels, func(tagRel *db.TagObjRel) int64 {
		return tagRel.TagId
	})

	modifyTagIdList := make([]int64, 0, len(tagList))
	for _, tagId := range tagIdList {
		if utils.NotContains(tagRelIdList, tagId) {
			modifyTagIdList = append(modifyTagIdList, tagId)
		}
	}

	tagObjRels := utils.MapStructList(modifyTagIdList, func(tagId int64) *db.TagObjRel {
		return &db.TagObjRel{
			Id:      utils.GetId(),
			ObjId:   objId,
			TagId:   tagId,
			ObjType: int32(sparkruntime.TagObjType_Candidate),
		}
	})

	if err = db.AddTagRels(ctx, db.DB, tagObjRels); err != nil {
		return 0, err
	}

	return len(tagObjRels), nil
}

func UnbindTags(ctx context.Context, objId int64, tagIdList []int64) (int, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return 0, err
	}

	// 1. candidate must exist
	candidate, err := db.FindCandidateById(ctx, db.DB, objId)
	if err != nil {
		return 0, err
	}
	if candidate.UserId != userId {
		return 0, fmt.Errorf("only can modify own tag")
	}

	if len(tagIdList) == 0 {
		return 0, nil
	}

	// 2. query existing tag relations
	tagRels, err := db.FindTagRelsByObjId(ctx, db.DB, objId, int32(sparkruntime.TagObjType_Candidate))
	if err != nil {
		return 0, err
	}

	// 已绑定的 tagId -> relId 映射
	tagId2RelId := make(map[int64]int64, len(tagRels))
	for _, rel := range tagRels {
		tagId2RelId[rel.TagId] = rel.Id
	}

	// 3. 过滤出真正需要解绑的 relId
	relIdList := make([]int64, 0, len(tagIdList))
	for _, tagId := range tagIdList {
		if relId, ok := tagId2RelId[tagId]; ok {
			relIdList = append(relIdList, relId)
		}
	}

	if len(relIdList) == 0 {
		return 0, nil
	}

	// 4. delete relations
	if err := db.DeleteTagRels(ctx, db.DB, relIdList); err != nil {
		return 0, err
	}

	return len(relIdList), nil
}
