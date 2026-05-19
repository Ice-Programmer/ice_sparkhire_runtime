package milvus

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/utils"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

type MilvusField struct {
	Name        string
	FieldType   entity.FieldType
	Dim         int64
	MaxLength   int64
	IsPrimary   bool
	IsAutoID    bool
	Description string
	DataType    entity.FieldType
	MaxCapacity int64
}

type MilvusSchema []MilvusField

func (s MilvusSchema) GenCollectionSchema(ctx context.Context, collectionName string) *entity.Schema {
	schema := entity.NewSchema().WithName(collectionName).WithDescription(collectionName)
	for _, field := range s {
		fieldSchema := entity.NewField().WithName(field.Name).WithDataType(field.FieldType)
		if len(field.Description) > 0 {
			fieldSchema = fieldSchema.WithDescription(field.Description)
		} else {
			fieldSchema = fieldSchema.WithDescription(field.Name)
		}
		if field.IsPrimary {
			fieldSchema = fieldSchema.WithIsPrimaryKey(true)
		}
		if field.IsAutoID {
			fieldSchema = fieldSchema.WithIsAutoID(true)
		}
		if field.FieldType == entity.FieldTypeVarChar {
			fieldSchema = fieldSchema.WithMaxLength(field.MaxLength)
		}
		if field.MaxLength > 0 {
			fieldSchema = fieldSchema.WithMaxLength(field.MaxLength)
		}
		if field.MaxCapacity > 0 {
			fieldSchema = fieldSchema.WithMaxCapacity(field.MaxCapacity)
		}
		if field.FieldType == entity.FieldTypeFloatVector {
			fieldSchema = fieldSchema.WithDim(field.Dim)
		}
		if field.FieldType == entity.FieldTypeArray {
			fieldSchema = fieldSchema.WithDataType(field.DataType)
		}
		schema.WithField(fieldSchema)
	}
	klog.CtxInfof(ctx, "Generate schema for collection %s, schema: %s", collectionName, utils.MarshalString(schema))
	return schema
}

func (s MilvusSchema) GetFirstVectorColName() string {
	for _, field := range s {
		if field.FieldType == entity.FieldTypeFloatVector {
			return field.Name
		}
	}
	return ""
}

func (s MilvusSchema) GetColNameList() []string {
	return utils.MapStructList(s, func(f MilvusField) string {
		return f.Name
	})
}

func (s MilvusSchema) GetPrimaryColumnName() string {
	for _, field := range s {
		if field.IsPrimary {
			return field.Name
		}
	}
	return ""
}
func (s MilvusSchema) GetNameAndTypeMap(fieldNames ...string) map[string]entity.FieldType {
	if len(fieldNames) == 0 {
		return utils.ToMap(s,
			func(schema MilvusField) string { return schema.Name },
			func(schema MilvusField) entity.FieldType { return schema.FieldType },
		)
	}

	schemaMap := make(map[string]entity.FieldType, len(fieldNames))
	for _, field := range s {
		if utils.Contains(fieldNames, field.Name) {
			schemaMap[field.Name] = field.FieldType
		}
	}

	return schemaMap
}
