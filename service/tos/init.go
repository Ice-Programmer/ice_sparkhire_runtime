package tos

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"ice_sparkhire_runtime/consts"
	"os"
)

var (
	tosClient *tos.ClientV2
	UrlPrefix string
)

func InitTos(ctx context.Context) error {
	var (
		ak       = os.Getenv("TOS_ACCESS_KEY")
		sk       = os.Getenv("TOS_SECRET_KEY")
		endpoint = os.Getenv("TOS_ENDPOINT")
		region   = os.Getenv("TOS_REGION")
	)
	UrlPrefix = fmt.Sprintf("https://%s.%s/", consts.TosBucketName, endpoint)
	credential := tos.NewStaticCredentials(ak, sk)
	client, err := tos.NewClientV2(endpoint, tos.WithCredentials(credential), tos.WithRegion(region))
	if err != nil {
		klog.CtxErrorf(ctx, "[TOS] init tos error: %s", err)
		return err
	}
	tosClient = client

	return nil
}
