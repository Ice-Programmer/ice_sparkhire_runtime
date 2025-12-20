package major

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListMajor(ctx context.Context) ([]*sparkruntime.MajorInfo, error) {
	majorList, err := db.ListAllMajor(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	majorInfoList := utils.MapStructList(majorList, func(major *db.Major) *sparkruntime.MajorInfo {
		return &sparkruntime.MajorInfo{
			Id:        major.Id,
			MajorName: major.MajorName,
		}
	})

	return majorInfoList, nil
}
