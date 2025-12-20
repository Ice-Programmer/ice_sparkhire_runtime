package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/utils"
	"sync"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// AuthInterceptor 把 base 中的 token 放入 ctx 中
func AuthInterceptor() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {
			buf := bufPool.Get().(*bytes.Buffer)
			buf.Reset()
			defer bufPool.Put(buf)

			if err := Obj2JsonBuf(req, buf); err == nil {
				var r Request
				if err := json.Unmarshal(buf.Bytes(), &r); err == nil {
					if token, ok := r.Req.Base.Extra[consts.AuthorizationHeader]; ok {
						ctx = utils.ContextSetKeyValue(ctx, consts.AuthorizationHeader, token)
					}
				}
			}
			return next(ctx, req, resp)
		}
	}
}

func Obj2JsonBuf(obj interface{}, buf *bytes.Buffer) error {
	buf.Reset()
	encoder := json.NewEncoder(buf)
	return encoder.Encode(obj)
}

type Request struct {
	Req struct {
		Base struct {
			Extra map[string]string `json:"Extra"`
		} `json:"Base"`
	} `json:"req"`
}
