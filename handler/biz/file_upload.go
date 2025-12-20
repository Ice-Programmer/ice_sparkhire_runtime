package biz

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
)

func UploadFile(ctx context.Context, req *sparkruntime.UploadFileRequest) (*sparkruntime.UploadFileResponse, error) {
	base := req.GetBase()

	marshal, err := sonic.Marshal(base)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	return &sparkruntime.UploadFileResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
