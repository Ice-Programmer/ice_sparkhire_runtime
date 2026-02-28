package wish_career

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/service/wish_career"
	"ice_sparkhire_runtime/utils"
)

func ModifyWishCareer(ctx context.Context, req *sparkruntime.ModifyWishCareerRequest) (*sparkruntime.ModifyWishCareerResponse, error) {
	currentUser, err := user.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if err := validateModifyWishCareer(ctx, req, currentUser); err != nil {
		return nil, err
	}

	if req.IsSetId() {
		err := db.UpdateWishCareerById(ctx, db.DB, req.GetId(), buildUpdateMap(req))
		if err != nil {
			return nil, err
		}
	} else {
		if err := db.CreateWishCareer(ctx, db.DB, buildWishCareer(req, currentUser.Id)); err != nil {
			return nil, err
		}

	}

	return &sparkruntime.ModifyWishCareerResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateModifyWishCareer(ctx context.Context, req *sparkruntime.ModifyWishCareerRequest, currentUser *db.User) error {
	if req.IsSetId() {
		if req.GetId() <= 0 {
			return fmt.Errorf("id is invalid")
		}

		wishCareer, err := db.FindWishCareerById(ctx, db.DB, req.GetId())
		if err != nil {
			return fmt.Errorf("find wish career: %w", err)
		}

		if wishCareer.UserId != currentUser.Id {
			return fmt.Errorf("only can Modify own wish career")
		}

		if wishCareer.CareerId != req.CareerId {
			if err := wish_career.EnsureWishCareerNotExists(ctx, currentUser.Id, req.CareerId); err != nil {
				return err
			}
		}
	} else {
		if sparkruntime.UserRole(currentUser.UserRole) != sparkruntime.UserRole_Candidate {
			return fmt.Errorf("user is not candidate")
		}

		if err := wish_career.EnsureWishCareerNotExists(ctx, currentUser.Id, req.CareerId); err != nil {
			return err
		}
	}

	if req.GetSalaryUpper() > 0 && req.GetSalaryLower() > 8 && req.GetSalaryLower() > req.GetSalaryUpper() {
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

func buildUpdateMap(req *sparkruntime.ModifyWishCareerRequest) map[string]interface{} {
	updateMap := map[string]interface{}{
		"id":             req.Id,
		"currency_type":  req.GetCurrencyType(),
		"frequency_type": req.GetFrequencyType(),
	}

	if req.IsSetSalaryLower() {
		updateMap["salary_lower"] = req.GetSalaryLower()
	}

	if req.IsSetSalaryUpper() {
		updateMap["salary_upper"] = req.GetSalaryUpper()
	}

	return updateMap
}

func buildWishCareer(req *sparkruntime.ModifyWishCareerRequest, userId int64) *db.CandidateWishCareer {
	wc := &db.CandidateWishCareer{
		Id:            utils.GetId(),
		UserId:        userId,
		CareerId:      req.GetCareerId(),
		CurrencyType:  int32(req.GetCurrencyType()),
		FrequencyType: int8(req.GetFrequencyType()),
	}

	utils.ApplyOptionalValue(req.IsSetSalaryLower, req.GetSalaryLower, &wc.SalaryLower)
	utils.ApplyOptionalValue(req.IsSetSalaryUpper, req.GetSalaryUpper, &wc.SalaryUpper)
	return wc
}
