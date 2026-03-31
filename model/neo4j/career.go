package neo4j

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CareerNode struct {
	ID         int64  `json:"id"`
	CareerName string `json:"career_name"`
}

func InsertCareerNode(ctx context.Context, node *CareerNode) error {
	params := map[string]any{
		"id":          node.ID,
		"career_name": node.CareerName,
	}
	query := "MERGE (t:Career {career_name: $career_name}) SET t.id = $id"
	_, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Neo4jDatabase),
	)
	if err != nil {
		klog.CtxErrorf(ctx, "[CareerNode][InsertData] error: %v", err)
		return err
	}

	return nil
}
