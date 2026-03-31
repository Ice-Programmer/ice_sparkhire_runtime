package neo4j

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TagNode struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
}

func InsertTagNode(ctx context.Context, node *TagNode) error {
	params := map[string]any{
		"id":       node.ID,
		"tag_name": node.TagName,
	}
	query := "MERGE (t:Tag {tag_name: t.tag_name}) SET t.id = t.id"
	_, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Neo4jDatabase),
	)
	if err != nil {
		klog.CtxErrorf(ctx, "[TagNode][InsertData] error: %v", err)
		return err
	}
	return nil
}

func BatchInsertTagNode(ctx context.Context, nodeList []*TagNode) error {
	if len(nodeList) == 0 {
		return nil
	}

	batchParams := make([]map[string]any, 0, len(nodeList))
	for _, node := range nodeList {
		batchParams = append(batchParams, map[string]any{
			"id":       node.ID,
			"tag_name": node.TagName,
		})
	}

	query := `
		UNWIND $rows AS row
        MERGE (t:Tag {tag_name: row.tag_name})
        SET t.id = row.id
    `
	_, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		map[string]any{"rows": batchParams},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Neo4jDatabase),
	)

	if err != nil {
		klog.CtxErrorf(ctx, "[TagNode][BatchInsert] error: %v", err)
		return err
	}

	return nil
}
