package career

import (
	"context"
	"fmt"
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
