package utils

import (
	"fmt"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	if len(email) > 254 {
		return fmt.Errorf("email is too long")
	}

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("email is invalid")
	}

	return nil
}

func ValidateVerifyCode(verifyCode string) error {
	verifyCode = strings.TrimSpace(verifyCode)

	if len(verifyCode) == 0 {
		return fmt.Errorf("verify code is empty")
	}

	if len(verifyCode) != 6 {
		return fmt.Errorf("verify code is invalid")
	}

	return nil
}

func ValidateYear(year int32) error {
	if year < 1900 || year > 2099 {
		return fmt.Errorf("year is invalid")
	}

	return nil
}

func ValidateBeginYearAndEndYear(beginYear int32, endYear int32) error {
	if err := ValidateYear(beginYear); err != nil {
		return err
	}

	if err := ValidateYear(endYear); err != nil {
		return err
	}

	if beginYear > endYear {
		return fmt.Errorf("begin year can not later than end year")
	}

	return nil
}

func ValidateLatitudeAndLongitude(latitude, longitude float64) error {
	if latitude < -90 || latitude > 90 {
		return fmt.Errorf("latitude is invalid")
	}

	if longitude < -180 || longitude > 180 {
		return fmt.Errorf("longitude is invalid")
	}

	return nil
}

func ValidateGeoDetail(geo *sparkruntime.GeoDetailInfo) error {
	if geo == nil {
		return fmt.Errorf("geo is nil")
	}

	if geo.FirstGeoLevelId <= 0 || geo.SecondGeoLevelId <= 0 || geo.ThirdGeoLevelId <= 0 || geo.ForthGeoLevelId <= 0 {
		return fmt.Errorf("geo is invalid")
	}

	if len(geo.GetAddress()) > 200 {
		return fmt.Errorf("address is too long")
	}

	return ValidateLatitudeAndLongitude(geo.GetLatitude(), geo.GetLongitude())
}
