package gateway

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/consts"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
	"strings"
)

const (
	CandidatePath = "/user/candidate/"
	HrPath        = "/user/hr/"
)

func CheckUserAuth(token, path string) (errorCode int32, err error) {
	for _, whitePath := range sparkruntime.UnAuthPathList {
		if strings.Contains(path, whitePath) {
			return 0, nil
		}
	}

	parseToken, err := utils.ParseToken(token)
	if err != nil {
		klog.Errorf("token parse error: %s", err.Error())
		return consts.InvalidTokenErrorCode, fmt.Errorf("token parse error, please login again")
	}

	// candidate
	if strings.Contains(path, CandidatePath) {
		user, err := db.FindUserById(context.Background(), db.DB, parseToken.ID)
		if err != nil {
			return consts.NoAuthErrorCode, fmt.Errorf("find user error: %s", err.Error())
		}

		if sparkruntime.UserRole(user.UserRole) != sparkruntime.UserRole_Candidate {
			return consts.NoAuthErrorCode, fmt.Errorf("only candidate role can operate")
		}
	}

	// hr
	if strings.Contains(path, HrPath) {
		user, err := db.FindUserById(context.Background(), db.DB, parseToken.ID)
		if err != nil {
			return consts.NoAuthErrorCode, fmt.Errorf("find user error: %s", err.Error())
		}

		if sparkruntime.UserRole(user.UserRole) != sparkruntime.UserRole_HR {
			return consts.NoAuthErrorCode, fmt.Errorf("only hr role can operate")
		}
	}

	return 0, nil
}
