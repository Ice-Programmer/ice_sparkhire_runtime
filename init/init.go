package init

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/mail"
	"ice_sparkhire_runtime/service/redis"
	"ice_sparkhire_runtime/service/tos"
	"ice_sparkhire_runtime/utils"
)

func Init(ctx context.Context) error {
	// init db
	klog.CtxInfof(ctx, "start init db")
	if err := db.InitDBGorm(); err != nil {
		klog.CtxErrorf(ctx, "init db err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "init db successfully!")

	// init redis
	klog.CtxInfof(ctx, "start init redis")
	if err := redis.InitRedis(ctx); err != nil {
		klog.CtxErrorf(ctx, "init redis err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "init redis successfully!")

	// init id generator
	klog.CtxInfof(ctx, "start init id generator")
	if err := utils.InitIdGeneratorClient(); err != nil {
		klog.CtxErrorf(ctx, "init id generator err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "init id generator successfully!")

	// init mail
	klog.CtxInfof(ctx, "start init mail")
	if err := mail.Init(ctx); err != nil {
		klog.CtxErrorf(ctx, "init mail err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "init mail successfully!")

	// init tos
	klog.CtxInfof(ctx, "start init tos")
	if err := tos.InitTos(ctx); err != nil {
		klog.CtxErrorf(ctx, "init tos err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "init tos successfully!")

	return nil
}
