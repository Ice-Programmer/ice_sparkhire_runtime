package candidate

import (
	"context"
	"fmt"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func BuildCandidateInfo(ctx context.Context, candidate *db.Candidate) (*sparkruntime.CandidateInfo, error) {
	geoInfo, err := geo.BuildGeoDetailInfo(ctx, candidate.FirstGeoLevelId, candidate.SecondGeoLevelId, candidate.ThirdGeoLevelId, candidate.ForthGeoLevelId, candidate.Address)
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

	currentUser, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.CandidateInfo{
		Age:       candidate.Age,
		Profile:   candidate.Profile,
		JobStatus: sparkruntime.JobStatus(candidate.JobStatus),
		ContractInfo: &sparkruntime.ContractInfo{
			Phone:   candidate.Phone,
			Email:   currentUser.Email,
			GeoInfo: geoInfo,
		},
		GraduationYear:  candidate.GraduationYear,
		EducationStatus: sparkruntime.EducationStatus(candidate.Education),
		Id:              utils.Int64Ptr(candidate.Id),
		TagList:         tagInfoList,
	}, nil
}

func BindTags(ctx context.Context, objId int64, tagIdList []int64) (int, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return 0, err
	}

	if objId != userId {
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
	if objId != userId {
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
