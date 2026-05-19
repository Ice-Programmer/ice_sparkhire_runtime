package milvus

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var (
	milvusClient client.Client
)

func InitMilvusClient(ctx context.Context) error {
	newClient, err := client.NewClient(ctx, client.Config{
		Address: "127.0.0.1:19530",
	})
	if err != nil {
		return err
	}

	milvusClient = newClient

	return nil
}
