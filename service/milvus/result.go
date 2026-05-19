package milvus

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/utils"
	"reflect"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

const MaxScoreThreshold = 1

const ScoreCol = "score"

func ParseSearchResult[T any](ctx context.Context, searchResults []client.SearchResult, schema *MilvusSchema, returnCols ...string) ([]*T, error) {
	if len(searchResults) == 0 {
		return nil, fmt.Errorf("no search results")
	}

	searchResult := searchResults[0]
	return ParseResultSet[T](ctx, searchResult.Fields, schema, returnCols, searchResult.Scores)
}

func ParseResultSet[T any](ctx context.Context, resultSet client.ResultSet, schema *MilvusSchema, returnCols []string, scoreValues []float32) ([]*T, error) {
	var zero T
	if reflect.TypeOf(zero).Kind() == reflect.Ptr {
		return nil, fmt.Errorf("generic type T must be a non-pointer struct type, got pointer")
	}

	schemaMap := schema.GetNameAndTypeMap(returnCols...)
	klog.CtxInfof(ctx, "[Milvus Parse] return Schema: %s", schemaMap)

	rows := make([]map[string]interface{}, 0)
	for i := 0; i < resultSet.Len(); i++ {
		row := make(map[string]interface{})
		if len(scoreValues) > 0 {
			if i < len(scoreValues) {
				row[ScoreCol] = scoreValues[i]
			} else {
				continue
			}
		}

		for columnName, columnType := range schemaMap {
			originValue := resultSet.GetColumn(columnName)
			commonValue, err := ConvertCommonType(originValue, columnType, i)
			if err != nil {
				return nil, err
			}
			row[columnName] = commonValue
		}

		rows = append(rows, row)
	}

	klog.CtxInfof(ctx, "[Milvus Parse] return rows: %s", utils.MarshalString(rows))

	result := make([]*T, 0, len(rows))
	for idx, r := range rows {
		b, err := json.Marshal(r)
		if err != nil {
			return nil, fmt.Errorf("marshal row %d failed: %v", idx, err)
		}

		var t T
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, fmt.Errorf("unmarshal row %d into T failed: %v; json=%s", idx, err, string(b))
		}
		result = append(result, &t)
	}

	return result, nil
}

func ConvertCommonType(milvusColumn entity.Column, milvusType entity.FieldType, index int) (commonValue interface{}, err error) {
	switch milvusType {
	case entity.FieldTypeVarChar, entity.FieldTypeString:
		commonValue, err = milvusColumn.GetAsString(index)
	case entity.FieldTypeInt64:
		commonValue, err = milvusColumn.GetAsInt64(index)
	case entity.FieldTypeFloat, entity.FieldTypeDouble:
		commonValue, err = milvusColumn.GetAsDouble(index)
	case entity.FieldTypeFloatVector:
		commonValue, err = milvusColumn.Get(index)
	case entity.FieldTypeBool:
		commonValue, err = milvusColumn.GetAsBool(index)
	default:
		commonValue, err = milvusColumn.Get(index)
	}

	if err != nil {
		return nil, err
	}
	return commonValue, nil
}
