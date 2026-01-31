package school

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListSchool(ctx context.Context) ([]*sparkruntime.SchoolInfo, error) {
	schoolList, err := db.ListAllSchool(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	schoolInfoList := utils.MapStructList(schoolList, func(school *db.School) *sparkruntime.SchoolInfo {
		return &sparkruntime.SchoolInfo{
			Id:         school.Id,
			SchoolName: school.SchoolName,
			SchoolIcon: school.SchoolIcon,
		}
	})

	return schoolInfoList, nil
}
