package elasticsearch

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v9"
)

var (
	elasticClient *elasticsearch.TypedClient
)

func InitElasticSearchClient() error {
	client, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("http://localhost:9200"),
		elasticsearch.WithAPIKey("ZEVqME9wNEIyZ0JmdlhFbWlpaUk6cW43dHN2OHl5Z0hMZ0ppZkFndVc2UQ=="),
	)
	if err != nil {
		return err
	}

	if client == nil {
		return errors.New("elasticsearch client is nil")
	}

	elasticClient = client

	return nil
}
