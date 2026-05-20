package message

import (
	"context"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
	"time"
)

func CreateChatMessage(ctx context.Context, req *sparkruntime.CreateChatMessageRequest) (*sparkruntime.CreateChatMessageResponse, error) {
	session, err := db.FindChatSessionById(ctx, db.DB, req.GetSessionId())
	if err != nil {
		return nil, err
	}

	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		if err := db.CreateChatMessage(ctx, tx, &db.ChatMessage{
			Id:           utils.GetId(),
			SessionId:    session.Id,
			SenderId:     userId,
			ReceiverId:   session.HrUserId,
			SenderType:   sparkruntime.SenderType_Candidate,
			MessageType:  sparkruntime.MessageType_Text,
			Content:      req.GetContent(),
			IsRead:       1,
			RevokeStatus: 0,
			SendStatus:   2,
			CreateTime:   time.Now(),
		}); err != nil {
			return err
		}

		if err := db.EditChatSessionLastMessage(ctx, tx, session.Id, req.Content); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &sparkruntime.CreateChatMessageResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
