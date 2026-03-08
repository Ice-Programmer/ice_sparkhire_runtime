package company

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/model/db"
)

func ValidateCompanyId(ctx context.Context, companyId int64) error {
	if companyId <= 0 {
		return fmt.Errorf("invalid company id")
	}

	if _, err := db.FindCompanyById(ctx, db.DB, companyId); err != nil {
		return err
	}

	return nil
}
