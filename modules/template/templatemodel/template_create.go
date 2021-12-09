package templatemodel

import (
	"lift-tracker-api/common"
	"strings"

	"gorm.io/datatypes"
)

type TemplateCreate struct {
	common.SimpleSQLModel `json:",inline"`
	Name                  string         `json:"name" form:"name" binding:"required" gorm:"column:name;"`
	CreatedBy             int            `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy         string         `json:"created_by" form:"created_by" binding:"required" gorm:"-"`
	Detail                datatypes.JSON `json:"detail" form:"detail" binding:"required" gorm:"column:detail;"`
	CollectionId          int            `json:"-" gorm:"column:collection_id;"`
	FakeCollectionId      string         `json:"collection_id" form:"collection_id" gorm:"-"`
	Note                  string         `json:"note" form:"note" gorm:"column:note;"`
	IsFavorite            bool           `json:"is_favorite" form:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived            bool           `json:"is_archived" form:"is_archived" gorm:"column:is_archived;"`
}

func (TemplateCreate) TableName() string {
	return Template{}.TableName()
}

func (data *TemplateCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeTemplate)
}

func (res *TemplateCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.FakeCreatedBy = strings.TrimSpace(res.FakeCreatedBy)
	res.FakeCollectionId = strings.TrimSpace(res.FakeCollectionId)
	res.Note = strings.TrimSpace(res.Note)

	if res.FakeCollectionId != "" {
		collectionId, err := common.FromBase58(strings.TrimSpace(res.FakeCollectionId))
		if err != nil {
			return err
		}

		res.CreatedBy = int(collectionId.GetLocalID())
	}

	if res.FakeCreatedBy != "" {
		createdBy, err := common.FromBase58(strings.TrimSpace(res.FakeCreatedBy))
		if err != nil {
			return err
		}

		res.CreatedBy = int(createdBy.GetLocalID())
	}

	return nil
}
