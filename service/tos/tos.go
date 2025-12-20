package tos

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"time"
)

func UploadObject(ctx context.Context, fileKey string, fileBytes []byte) error {
	if _, err := tosClient.PutObjectV2(ctx, &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: "ice-spark-hire",
			Key:    fileKey,
		},
		Content: bytes.NewReader(fileBytes),
	}); err != nil {
		klog.CtxErrorf(ctx, "[TOS] Failed to upload object: %v", err)
		return err
	}
	return nil
}

func BuildFileKey(fileName string, fileType sparkruntime.FileType) string {
	return fmt.Sprintf("%s/%d_%s", sparkruntime.FileTypeMap[fileType], time.Now().Unix(), fileName)
}
