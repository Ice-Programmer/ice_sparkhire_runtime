package biz

import (
	"context"
	"encoding/base64"
	"fmt"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/service/tos"
)

func UploadFile(ctx context.Context, req *sparkruntime.UploadFileRequest) (*sparkruntime.UploadFileResponse, error) {
	if len(req.GetFile()) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	if len(req.GetFilename()) == 0 {
		return nil, fmt.Errorf("filename is empty")
	}

	rawData := req.File
	if isBase64Content(rawData) {
		decoded, err := base64.StdEncoding.DecodeString(string(rawData))
		if err == nil {
			rawData = decoded
		}
	}

	fileKey := tos.BuildFileKey(req.Filename, req.GetFileType())
	if err := tos.UploadObject(ctx, fileKey, rawData); err != nil {
		return nil, err
	}

	return &sparkruntime.UploadFileResponse{
		FileLink: tos.BuildFileURL(fileKey),
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func isBase64Content(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	return true
}
