package utils

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
)

func GetCurrentUserId(ctx context.Context) (int64, error) {
	value := ContextGetKeyValue(ctx, consts.AuthorizationHeader)
	if value == nil {
		return 0, fmt.Errorf("user not login")
	}

	parseToken, err := ParseToken(value.(string))
	if err != nil || parseToken == nil {
		return 0, fmt.Errorf("parse token error: %v", err)
	}

	return parseToken.ID, nil
}
