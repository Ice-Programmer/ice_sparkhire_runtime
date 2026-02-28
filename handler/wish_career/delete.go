package wish_career

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func DeleteWishCareer(ctx context.Context, req *sparkruntime.DeleteWishCareerRequest) (*sparkruntime.DeleteWishCareerResponse, error) {
	if req.GetId() <= 0 {
		return nil, fmt.Errorf("id is invalid")
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	wishCareer, err := db.FindWishCareerById(ctx, db.DB, req.GetId())
	if err != nil {
		return nil, err
	}

	if wishCareer.UserId != userId {
		return nil, fmt.Errorf("only can delete own wish career")
	}

	if err := db.DeleteWishCareerById(ctx, db.DB, req.GetId()); err != nil {
		return nil, err
	}

	return &sparkruntime.DeleteWishCareerResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
