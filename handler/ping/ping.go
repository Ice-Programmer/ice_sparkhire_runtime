package ping

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/utils"
)

func Ping(ctx context.Context, req *sparkruntime.PingRequest) (*sparkruntime.PingResponse, error) {
	klog.CtxInfof(ctx, "received ping request: %v", req)

	if len(req.GetPing()) == 0 {
		req.SetPing(utils.StringPtr("Ping"))
	}
	return &sparkruntime.PingResponse{
		Pong:     req.GetPing() + " Pong",
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
