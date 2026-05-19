package model

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"

	"os"
)

var (
	apiKey  string
	model   string
	baseURL string
)

func init() {
	apiKey = os.Getenv("DEEPSEEK_API_KEY")
	model = os.Getenv("DEEPSEEK_MODEL")
	baseURL = os.Getenv("ARK_URL")
}

func NewCommonChatModel(ctx context.Context) (*openai.ChatModel, error) {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		Model:   model,
		BaseURL: baseURL,
	})

	if err != nil {
		return nil, err
	}

	return chatModel, nil
}

func NewEmbeddingModel(ctx context.Context) (*openai.ChatModel, error) {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		Model:   "doubao-embedding-vision-250615",
		BaseURL: baseURL,
	})
	if err != nil {
		return nil, err
	}

	return chatModel, nil
}
