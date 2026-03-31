package neo4j

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	CareerToTagRelation = "HAS_TAG"
)

func BindCareerToTag(ctx context.Context, careerName, tagName string) error {
	query := fmt.Sprintf(`
       MATCH (c:Career {career_name: $career_name})
       MATCH (t:Tag {tag_name: $tag_name})
       MERGE (c)-[r:%s]->(t)
       RETURN type(r)
    `, CareerToTagRelation)
	params := map[string]any{
		"career_name": careerName,
		"tag_name":    tagName,
	}

	_, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Neo4jDatabase),
	)
	if err != nil {
		klog.CtxErrorf(ctx, "[Relationship][BindCareerToTag] error: %v", err)
		return err
	}
	return nil
}

func BatchBindCareerToTags(ctx context.Context, careerName string, tagNames []string) error {
	query := fmt.Sprintf(`
       MATCH (c:Career {career_name: $career_name})
       UNWIND $tag_list AS t_name
       MATCH (t:Tag {tag_name: t_name})
       MERGE (c)-[:%s]->(t)
    `, CareerToTagRelation)

	params := map[string]any{
		"career_name": careerName,
		"tag_list":    tagNames,
	}

	_, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Neo4jDatabase),
	)
	return err
}
