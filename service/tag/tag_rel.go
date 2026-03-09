package tag

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
)

func GetTagInfoMap(ctx context.Context, objs []int64) (map[int64]*sparkruntime.TagInfo, error) {
	db.FindTagByName()
}
