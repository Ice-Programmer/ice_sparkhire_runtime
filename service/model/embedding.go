package model

import (
	"bytes"
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

type TextInput struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type EmbeddingRequest struct {
	Model          string      `json:"model"`
	EncodingFormat string      `json:"encoding_format"`
	Input          []TextInput `json:"input"`
}

type EmbeddingResponse struct {
	Data struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

func EmbeddingText(ctx context.Context, text string) ([]float32, error) {
	reqBody := EmbeddingRequest{
		Model:          "doubao-embedding-vision-250615",
		EncodingFormat: "float",
		Input: []TextInput{
			{
				Type: "text",
				Text: text,
			},
		},
	}

	jsonData, err := sonic.Marshal(reqBody)
	req, err := http.NewRequest(
		"POST",
		"https://ark.cn-beijing.volces.com/api/v3/embeddings/multimodal",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var embeddingResponse EmbeddingResponse
	if err := sonic.Unmarshal(body, &embeddingResponse); err != nil {
		panic(err)
	}

	if len(embeddingResponse.Data.Embedding) == 0 {
		return nil, errors.New("no embedding data")
	}

	return embeddingResponse.Data.Embedding, nil
}
