package chat

import (
	"context"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/user"
	"ice_sparkhire_runtime/utils"
)

func BuildChatMessageInfoList(ctx context.Context, messageList []*db.ChatMessage) ([]*sparkruntime.ChatMessageInfo, error) {
	if len(messageList) == 0 {
		return []*sparkruntime.ChatMessageInfo{}, nil
	}

	senderIdList := utils.MapStructListDistinct(messageList, func(message *db.ChatMessage) int64 {
		return message.SenderId
	})

	userBasicInfoMap, err := user.GetUserBasicInfoMap(ctx, senderIdList)
	if err != nil {
		return nil, err
	}

	chatMessageInfoList := utils.MapStructList(messageList, func(message *db.ChatMessage) *sparkruntime.ChatMessageInfo {
		return &sparkruntime.ChatMessageInfo{
			Id:          message.Id,
			Content:     message.Content,
			SenderInfo:  userBasicInfoMap[message.SenderId],
			MessageType: message.MessageType,
			IsRead:      message.IsRead == 1,
			CreatedAt:   message.CreateTime.Unix(),
		}
	})

	return chatMessageInfoList, nil
}
