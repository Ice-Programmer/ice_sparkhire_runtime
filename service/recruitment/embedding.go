package recruitment

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"ice_sparkhire_runtime/consts"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/service/milvus"
	"ice_sparkhire_runtime/service/model"
)

const RecruitmentCollectionName = "sparkhire_recruitment"

const (
	IdCol            = "id"
	VectorCol        = "vector"
	RecruitmentIdCol = "recruitment_id"

	RecruitmentDescCol = "recruitment_description"
)

type RecruitmentMilvusItem struct {
	RecruitmentDesc string  `json:"recruitment_description"`
	Score           float64 `json:"score"`
	RecruitmentID   int64   `json:"recruitment_id"`
	ID              int64   `json:"id"`
}

var RecruitmentCollectionSchema = &milvus.MilvusSchema{
	milvus.MilvusField{Name: IdCol, FieldType: entity.FieldTypeInt64, IsPrimary: true, IsAutoID: true},
	milvus.MilvusField{Name: VectorCol, FieldType: entity.FieldTypeFloatVector, Dim: consts.Dim},
	milvus.MilvusField{Name: RecruitmentIdCol, FieldType: entity.FieldTypeInt64},
	milvus.MilvusField{Name: RecruitmentDescCol, FieldType: entity.FieldTypeVarChar, MaxLength: consts.VarcharMaxLength},
}

func EmbeddingRecruitmentInfoToMilvus(ctx context.Context, recruitmentItemList []*RecruitmentMilvusItem, isCreateNewTable bool) error {
	milvusManager := milvus.NewMilvusManager(ctx, RecruitmentCollectionSchema, RecruitmentCollectionName)
	// 1. judge whether create new table
	if isCreateNewTable {
		// 1.1 create table
		if err := milvusManager.CreateCollection(); err != nil {
			return err
		}

		// 1.2 create index
		if err := milvusManager.CreateHNSWIndex(consts.NHSWIndexM, consts.NHSWIndexEf); err != nil {
			return err
		}
	}

	// 3. Insert Collection Index
	for _, recruitment := range recruitmentItemList {
		embeddedTableInfo, err := model.EmbeddingText(ctx, recruitment.RecruitmentDesc)
		if err != nil {
			return err
		}

		if err = milvusManager.UpsertData(
			entity.NewColumnFloatVector(VectorCol, consts.Dim, [][]float32{embeddedTableInfo}),
			entity.NewColumnInt64(IdCol, []int64{recruitment.ID}),
			entity.NewColumnInt64(RecruitmentIdCol, []int64{recruitment.RecruitmentID}),
			entity.NewColumnVarChar(RecruitmentDescCol, []string{recruitment.RecruitmentDesc}),
		); err != nil {
			return err
		}

	}

	// 4. Flush & Load Collection
	if err := milvusManager.FlushAndLoad(); err != nil {
		return err
	}

	return nil
}

func BuildRecruitmentMilvusItem(recruitment *db.Recruitment) *RecruitmentMilvusItem {
	return &RecruitmentMilvusItem{
		RecruitmentDesc: recruitment.Description + "-" + recruitment.Requirement,
		RecruitmentID:   recruitment.ID,
	}
}
