package tag

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/candidate"
)

func UnbindTags(ctx context.Context, req *sparkruntime.UnbindTagsRequest) (resp *sparkruntime.UnbindTagsResponse, err error) {
	if len(req.GetTagIdList()) == 0 {
		return nil, fmt.Errorf("tag id list is empty")
	}

	if req.GetObjId() <= 0 {
		return nil, fmt.Errorf("obj id is invalid")
	}

	var num int
	switch req.GetObjType() {
	case sparkruntime.TagObjType_Candidate:
		num, err = candidate.UnbindTags(ctx, req.GetObjId(), req.GetTagIdList())
	case sparkruntime.TagObjType_Recruitment:
		// todo 补充逻辑
	default:
		return nil, fmt.Errorf("unsupported obj type: %s", req.GetObjType())
	}

	if err != nil {
		return nil, err
	}

	return &sparkruntime.UnbindTagsResponse{
		Num:      int64(num),
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
