package company

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func DeleteCompany(ctx context.Context, req *sparkruntime.DeleteCompanyRequest) (*sparkruntime.DeleteCompanyResponse, error) {
	if req.GetId() <= 0 {
		return nil, fmt.Errorf("invalid company id")
	}

	company, err := db.FindCompanyById(ctx, db.DB, req.GetId())
	if err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if company.CreateUserId != userId {
		return nil, fmt.Errorf("only creator can delete company")
	}

	if err = db.DeleteCompany(ctx, db.DB, req.GetId()); err != nil {
		return nil, err
	}

	return &sparkruntime.DeleteCompanyResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
