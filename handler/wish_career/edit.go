package wish_career

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/wish_career"
	"ice_sparkhire_runtime/utils"
)

func EditWishCareer(ctx context.Context, req *sparkruntime.EditWishCareerRequest) (*sparkruntime.EditWishCareerResponse, error) {
	if err := validateEditWishCareer(ctx, req); err != nil {
		return nil, err
	}

	err := db.UpdateWishCareerById(ctx, db.DB, req.Id, buildUpdateMap(req))
	if err != nil {
		return nil, err
	}

	return &sparkruntime.EditWishCareerResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateEditWishCareer(ctx context.Context, req *sparkruntime.EditWishCareerRequest) error {
	if req.GetId() <= 0 {
		return fmt.Errorf("id is invalid")
	}

	wishCareer, err := db.FindWishCareerById(ctx, db.DB, req.Id)
	if err != nil {
		return fmt.Errorf("find wish career: %w", err)
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return fmt.Errorf("get current user id: %w", err)
	}

	if wishCareer.UserId != userId {
		return fmt.Errorf("only can edit own wish career")
	}

	if wishCareer.CareerId != req.CareerId {
		if err := wish_career.EnsureWishCareerNotExists(ctx, userId, req.CareerId); err != nil {
			return err
		}
	}

	if req.IsSetSalaryLower() && req.IsSetSalaryUpper() && req.GetSalaryLower() > req.GetSalaryUpper() {
		return fmt.Errorf("lowest wish salary can not greater than highest salary")
	}

	if req.GetCareerId() <= 0 {
		return fmt.Errorf("career id is invalid")
	}

	if _, err := db.FindCareerById(ctx, db.DB, req.GetCareerId()); err != nil {
		return fmt.Errorf("find career error: %v", err)
	}

	if utils.NotContains(sparkruntime.SalaryCurrencyTypeList, req.GetCurrencyType()) {
		return fmt.Errorf("currency type is invalid")
	}

	if utils.NotContains(sparkruntime.SalaryFrequencyTypeList, req.GetFrequencyType()) {
		return fmt.Errorf("frequency type is invalid")
	}

	return nil
}

func buildUpdateMap(req *sparkruntime.EditWishCareerRequest) map[string]interface{} {
	updateMap := map[string]interface{}{
		"id":             req.Id,
		"currency_type":  req.GetCurrencyType(),
		"frequency_type": req.GetFrequencyType(),
	}

	if req.IsSetSalaryLower() && req.IsSetSalaryUpper() {
		updateMap["salary_lower"] = req.GetSalaryLower()
		updateMap["salary_upper"] = req.GetSalaryUpper()
	}

	return updateMap
}
