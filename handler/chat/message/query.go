package message

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/chat"
	"ice_sparkhire_runtime/utils"
)

func QueryChatMessage(ctx context.Context, req *sparkruntime.QueryChatMessageRequest) (*sparkruntime.QueryChatMessageResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.GetPageSize(), req.GetPageNum())

	messageList, total, err := db.QueryChatMessagePage(ctx, db.DB, pageSize, pageNum, req.GetSessionId())
	if err != nil {
		return nil, err
	}

	messageInfoList, err := chat.BuildChatMessageInfoList(ctx, messageList)
	if err != nil {
		return nil, err
	}

	return &sparkruntime.QueryChatMessageResponse{
		MessageList: messageInfoList,
		Total:       utils.Int64Ptr(total),
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
