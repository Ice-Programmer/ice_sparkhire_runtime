package user

import (
	"context"
	"fmt"
	"ice_sparkhire_runtime/consts"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/redis"
	"ice_sparkhire_runtime/utils"
	"math/rand"
	"strconv"
	"time"
)

func GetCurrentUser(ctx context.Context) (*db.User, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	userInfo, err := db.FindUserById(ctx, db.DB, userId)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func CompareVerifyCode(ctx context.Context, email, verifyCode string) error {
	verifyCodeKey := consts.VerifyCodePrefixKey + email
	if !redis.HasValue(ctx, verifyCodeKey) {
		return fmt.Errorf("verify code not exist")
	}

	verifyCacheValue, err := redis.GetValue(ctx, verifyCodeKey)
	if err != nil {
		return fmt.Errorf("verify code not exist")
	}

	if verifyCacheValue != verifyCode {
		return fmt.Errorf("verify code is wrong")
	}

	return nil
}

func BuildEmptyUser(ctx context.Context, email string) (*db.User, error) {
	user, err := db.FindUserByEmail(ctx, db.DB, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("email has been registered")
	}

	return &db.User{
		Id:         utils.GetId(),
		Username:   GenerateUniqueUsername(),
		UserAvatar: GetRandomUserAvatar(),
		Email:      email,
		UserRole:   int8(sparkruntime.UserRole_Visitor),
	}, nil
}

func BuildUserBasicInfo(user *db.User) *sparkruntime.UserBasicInfo {
	return &sparkruntime.UserBasicInfo{
		Id:         user.Id,
		Username:   user.Username,
		Role:       sparkruntime.UserRole(user.UserRole),
		UserAvatar: user.UserAvatar,
		Gender:     user.Gender,
		Email:      user.Email,
	}
}

func GenerateUniqueUsername() string {
	name := consts.DefaultUserList[rand.Intn(len(consts.DefaultUserList))]
	timestamp := time.Now().UnixMilli() % 1_000_000
	randomNum := rand.Intn(90) + 10
	return name + "#" + strconv.FormatInt(timestamp, 10) + strconv.Itoa(randomNum)
}

func GetRandomUserAvatar() string {
	return consts.DefaultAvatarList[rand.Intn(len(consts.DefaultAvatarList))]
}
