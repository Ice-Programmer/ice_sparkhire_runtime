package education_exp

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func DeleteEducationExp(ctx context.Context, req *sparkruntime.DeleteEducationExpRequest) (*sparkruntime.DeleteEducationExpResponse, error) {
	if req.GetId() <= 0 {
		return nil, fmt.Errorf("id is invalid")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := db.FindEducationExperienceByUserIdAndId(ctx, db.DB, req.GetId(), userId); err != nil {
		return nil, fmt.Errorf("find EducationExperience error: %v", err)
	}

	if err := db.DeleteEducationExperience(ctx, db.DB, req.GetId()); err != nil {
		return nil, fmt.Errorf("delete EducationExperience error: %v", err)
	}

	return &sparkruntime.DeleteEducationExpResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
