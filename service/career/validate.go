package career

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
)

func ValidateCareerId(ctx context.Context, careerId int64) error {
	if careerId <= 0 {
		return fmt.Errorf("invalid career id")
	}

	if _, err := db.FindCareerById(ctx, db.DB, careerId); err != nil {
		return err
	}

	return nil
}

func ValidateSalaryInfo(salaryInfo *sparkruntime.SalaryInfo) error {
	if salaryInfo == nil {
		return nil
	}

	if salaryInfo.GetSalaryUpper() > 0 && salaryInfo.GetSalaryLower() > 0 && salaryInfo.GetSalaryLower() > salaryInfo.GetSalaryUpper() {
		return fmt.Errorf("lowest wish salary can not greater than highest salary")
	}

	if salaryInfo.GetFrequencyType().String() == consts.EnumNotFound {
		return fmt.Errorf("frequency type is invalid")
	}

	if salaryInfo.GetCurrencyType().String() == consts.EnumNotFound {
		return fmt.Errorf("currency type is invalid")
	}

	return nil
}
