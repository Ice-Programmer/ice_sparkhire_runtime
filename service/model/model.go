package model

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"os"
)

func NewCommonChatModel(ctx context.Context) (*openai.ChatModel, error) {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  os.Getenv("DEEPSEEK_API_KEY"),
		Model:   os.Getenv("DEEPSEEK_MODEL"),
		BaseURL: os.Getenv("ARK_URL"),
	})

	if err != nil {
		return nil, err
	}

	return chatModel, nil
}
