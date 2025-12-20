package biz

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/mail"
	"ice_sparkhire_runtime/service/redis"
	"ice_sparkhire_runtime/utils"
	"os"
)

func SendVerifyCode(ctx context.Context, req *sparkruntime.SendVerifyCodeRequest) (resp *sparkruntime.SendVerifyCodeResponse, err error) {
	if err = utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}

	// 1. has sent code
	hasSendCodeKey := consts.HasSendCodePrefixKey + req.Email
	verifyCodeKey := consts.VerifyCodePrefixKey + req.Email

	if redis.HasValue(ctx, hasSendCodeKey) {
		return nil, fmt.Errorf("verify code has been sent, please try later")
	}

	// 2. generate verify code
	verifyCode := utils.Generate6DigitCode()
	klog.CtxInfof(ctx, "send verify code with verify code: %s", verifyCode)

	// 3. save in redis
	if err := redis.SetValue(ctx, hasSendCodeKey, consts.RecordValue, consts.HasSendVerifyCodeDuration); err != nil {
		return nil, err
	}
	if err := redis.SetValue(ctx, verifyCodeKey, verifyCode, consts.VerifyCodeDuration); err != nil {
		return nil, err
	}

	// 4. send mail
	// todo 改造成异步队列
	if err := mail.SendMail(
		ctx,
		os.Getenv("QQ_EMAIL_ADDRESS"),
		req.GetEmail(),
		mail.GenVerifyCodeContent(verifyCode),
		mail.VerifyCodeSubject,
	); err != nil {
		klog.CtxErrorf(ctx, "send verify code subject failed, err: %v", err)
		_ = redis.DelValue(ctx, hasSendCodeKey)
		_ = redis.DelValue(ctx, verifyCodeKey)
		return nil, err
	}

	return &sparkruntime.SendVerifyCodeResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
