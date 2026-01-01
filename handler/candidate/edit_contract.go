package candidate

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/geo"
	"ice_sparkhire_runtime/utils"
)

func UpsertCandidateContractInfo(ctx context.Context, req *sparkruntime.EditCandidateContractInfoRequest) (*sparkruntime.EditCandidateContractInfoResponse, error) {
	if err := validateContractInfo(ctx, req); err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := db.FindCandidateByUserId(ctx, db.DB, userId); err != nil {
		return nil, err
	}

	updateFields := map[string]interface{}{
		"first_geo_level_id":  req.GeoInfo.FirstGeoLevelId,
		"second_geo_level_id": req.GeoInfo.SecondGeoLevelId,
		"third_geo_level_id":  req.GeoInfo.ThirdGeoLevelId,
		"forth_geo_level_id":  req.GeoInfo.ForthGeoLevelId,
		"address":             req.GeoInfo.Address,
		"longitude":           req.GeoInfo.Longitude,
		"latitude":            req.GeoInfo.Latitude,
	}

	if req.IsSetPhoneNumber() {
		updateFields["phone"] = req.PhoneNumber
	}

	// edit contractInfo
	if err = db.UpdateCandidateByUserId(ctx, db.DB, userId, updateFields); err != nil {
		return nil, err
	}

	return &sparkruntime.EditCandidateContractInfoResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func validateContractInfo(ctx context.Context, req *sparkruntime.EditCandidateContractInfoRequest) error {
	if req.IsSetPhoneNumber() && len(req.GetPhoneNumber()) <= 3 {
		return fmt.Errorf("phone number is invalid")
	}

	if err := geo.ValidateGeoInfo(ctx, req.GetGeoInfo()); err != nil {
		return err
	}

	return nil
}
