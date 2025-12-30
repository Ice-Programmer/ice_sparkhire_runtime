package geo

import (
	"context"
	"fmt"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/utils"
)

func ValidateGeoInfo(ctx context.Context, geoInfo *sparkruntime.GeoModifyInfo) error {
	if len(geoInfo.GetAddress()) == 0 {
		return fmt.Errorf("address is empty")
	}

	if err := utils.ValidateLatitudeAndLongitude(geoInfo.GetLatitude(), geoInfo.GetLongitude()); err != nil {
		return err
	}

	return nil
}
