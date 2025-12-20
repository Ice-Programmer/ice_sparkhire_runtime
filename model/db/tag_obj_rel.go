package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

const (
	TagObjRelTableName = "tag_obj_rel"
)

type TagObjRel struct {
	Id        int64     `gorm:"column:id;type:bigint;comment:id;primaryKey;" json:"id"`                                             // id
	TagId     int64     `gorm:"column:tag_id;type:bigint;comment:tag id;not null;" json:"tag_id"`                                   // tag id
	ObjId     int64     `gorm:"column:obj_id;type:bigint;comment:obj_id;not null;" json:"obj_id"`                                   // obj_id
	ObjType   int32     `gorm:"column:obj_type;type:int;comment:obj type(1-candidate/2-recruitment);not null;" json:"obj_type"`     // obj type(1-candidate/2-recruitment)
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
}

func (TagObjRel) TableName() string {
	return TagObjRelTableName
}

func FindTagRelsByObjId(ctx context.Context, db *gorm.DB, objId int64, objType int32) ([]*TagObjRel, error) {
	var tagObjRels []*TagObjRel

	err := db.WithContext(ctx).
		Where("obj_id = ?", objId).
		Where("obj_type = ?", objType).
		Find(&tagObjRels).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] find tag rel error: %v", err)
		return nil, err
	}

	return tagObjRels, nil
}

func AddTagRels(ctx context.Context, db *gorm.DB, tagObjRels []*TagObjRel) error {
	if err := db.WithContext(ctx).Save(tagObjRels).Error; err != nil {
		klog.CtxErrorf(ctx, "[DB] add tag rel error: %v", err)
		return err
	}
	return nil
}

func DeleteTagRels(ctx context.Context, db *gorm.DB, tagRelIds []int64) error {
	if len(tagRelIds) == 0 {
		return nil
	}

	err := db.WithContext(ctx).Model(&TagObjRel{}).
		Where("id in (?)", tagRelIds).
		Delete(&TagObjRel{}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "[DB] delete tag rel error: %v", err)
		return err
	}
	return nil
}

func FindTagsByObjIdAndObjType(ctx context.Context, db *gorm.DB, objId int64, objType int32) ([]*Tag, error) {
	var tags []*Tag

	err := db.WithContext(ctx).
		Table("tag AS t").
		Select("t.*").
		Joins("JOIN tag_obj_rel AS r ON t.id = r.tag_id").
		Where("r.obj_type = ? AND r.obj_id = ?", objType, objId).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}
