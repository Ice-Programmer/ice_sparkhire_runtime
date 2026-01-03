package wish_career

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/model/db"
)

func EnsureWishCareerNotExists(ctx context.Context, userId, careerId int64) error {
	wishCareer, err := db.FindWishCareerByUserIdAndCareerId(ctx, db.DB, userId, careerId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if wishCareer != nil {
		return fmt.Errorf("wish career already exists")
	}
	return nil
}
