package neo4j

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	Neo4jDatabase = "neo4j"
)

var (
	driver neo4j.DriverWithContext
)

func InitNeo4j(ctx context.Context) error {
	var (
		user     = "neo4j"
		password = "12345678"
	)

	uri := "neo4j://localhost:7687"
	d, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to connect to Neo4j: %v", err)
		return err
	}

	err = d.VerifyConnectivity(ctx)
	if err != nil {
		return err
	}

	driver = d

	return nil
}
