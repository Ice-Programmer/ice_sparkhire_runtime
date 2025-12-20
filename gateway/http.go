package gateway

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/kitex_gen/base"
	"log"
	"net/http"
	"net/url"
)

func StartHttpServer() error {
	// 创建IDL Provider
	p, err := generic.NewThriftFileProvider(".conf/idl/ice_sparkhire_runtime.thrift", "./idl")
	if err != nil {
		log.Fatal("Failed to create IDL provider:", err)
		return err
	}

	// 创建HTTP映射的Generic Client
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		log.Fatal("Failed to create HTTP generic:", err)
		return err
	}

	genericCli, err := genericclient.NewClient("demo", g,
		client.WithHostPorts("127.0.0.1:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		log.Fatal("Failed to create generic client:", err)
		return err
	}

	// 创建HTTP服务器
	h := server.Default(server.WithHostPorts(":8000"))

	// 通用HTTP处理器，自动映射所有API
	h.Any("/*path", func(c context.Context, ctx *app.RequestContext) {
		// ===== 打印 HTTP 请求信息 =====
		log.Printf(
			"[Gateway Request] method=%s path=%s query=%s body=%s",
			string(ctx.Method()),
			string(ctx.URI().Path()),
			string(ctx.URI().QueryString()),
			string(ctx.Request.Body()),
		)

		// 创建标准的http.Request
		req := &http.Request{
			Method: string(ctx.Method()),
			Header: make(http.Header),
			Body:   nil,
		}

		// 设置URL
		uri := ctx.URI()
		path := string(uri.Path())
		req.URL = &url.URL{
			Path:     path,
			RawQuery: string(uri.QueryString()),
		}

		token := ctx.Request.Header.Get(consts.AuthorizationHeader)
		if errCode, err := CheckUserAuth(token, path); err != nil {
			ctx.JSON(http.StatusOK, base.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    errCode,
			})
			return
		}

		// 设置请求体
		// 复制请求头
		ctx.Request.Header.VisitAll(func(key, value []byte) {
			req.Header.Add(string(key), string(value))
		})

		// 如果上游没设置Content-Type，就默认用JSON
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}

		// 使用Kitex提供的FromHTTPRequest函数创建HTTPRequest
		httpReq, err := generic.FromHTTPRequest(req)
		if err != nil {
			log.Printf("Failed to create HTTPRequest: %v", err)
			ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		// 设置请求体数据
		httpReq.Body = make(map[string]interface{})
		if len(ctx.Request.Body()) > 0 {
			if err := json.Unmarshal(ctx.Request.Body(), &httpReq.Body); err != nil {
				log.Printf("Failed to unmarshal request body: %v", err)
				ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
				return
			}
		}
		baseMap := make(map[string]interface{})
		ctx.Request.Header.VisitAll(func(k, v []byte) {
			key := string(k)
			val := string(v)
			baseMap[key] = val
		})

		// 2️⃣ 将 Base 注入到请求 body 的 "Base" 字段
		if _, ok := httpReq.Body["Base"]; !ok {
			httpReq.Body["Base"] = map[string]interface{}{}
		}
		baseInterface, ok := httpReq.Body["Base"].(map[string]interface{})
		if !ok {
			baseInterface = map[string]interface{}{}
			httpReq.Body["Base"] = baseInterface
		}
		baseInterface["Extra"] = baseMap

		if bodyBytes, err := json.Marshal(httpReq.Body); err == nil {
			log.Printf(
				"[Gateway -> RPC] path=%s body=%s",
				string(uri.Path()),
				string(bodyBytes),
			)
		}

		bodyBytes, _ := json.Marshal(httpReq.Body)
		ctx.Request.AppendBodyString(string(bodyBytes))

		// 调用Generic Client
		resp, err := genericCli.GenericCall(c, "", httpReq)
		if err != nil {
			log.Printf("Generic call failed: %v", err)
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		if respBytes, err := json.Marshal(resp); err == nil {
			log.Printf(
				"[Gateway <- RPC] path=%s resp=%s",
				string(uri.Path()),
				string(respBytes),
			)
		}

		// 处理响应
		if httpResp, ok := resp.(*generic.HTTPResponse); ok {
			// 设置响应头
			for k, v := range httpResp.Header {
				if len(v) > 0 {
					ctx.Response.Header.Set(k, v[0])
				}
			}

			ctx.Response.Header.Set("Content-Type", "application/json")

			// 设置状态码
			ctx.Status(http.StatusOK)

			// 设置响应体
			if httpResp.Body != nil {
				respBytes, err := json.Marshal(httpResp.Body)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal response"})
					return
				}

				log.Printf(
					"[Gateway Response] path=%s body=%s",
					string(uri.Path()),
					string(respBytes),
				)

				ctx.Write(respBytes)
			}
		} else {
			// 如果不是HTTP响应格式，尝试JSON序列化
			respBytes, err := json.Marshal(resp)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal response"})
				return
			}
			ctx.JSON(http.StatusOK, json.RawMessage(respBytes))
		}
	})

	log.Println("HTTP Gateway server listening on :8000")
	log.Println("All HTTP requests will be automatically mapped to Thrift service methods")
	h.Spin()

	return nil
}
