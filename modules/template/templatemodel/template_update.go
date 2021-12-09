package templatemodel

import (
	"lift-tracker-api/common"
	"strings"

	"gorm.io/datatypes"
)

type TemplateUpdate struct {
	common.SimpleSQLModel `json:",inline"`
	Name                  string         `json:"name" form:"name" gorm:"column:name;"`
	CreatedBy             int            `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy         string         `json:"created_by" form:"created_by" gorm:"-"`
	Detail                datatypes.JSON `json:"detail" form:"detail" binding:"required" gorm:"column:detail;"`
	CollectionId          int            `json:"-" gorm:"column:collection_id;"`
	FakeCollectionId      string         `json:"collection_id" form:"collection_id" gorm:"-"`
	Note                  string         `json:"note" form:"note" gorm:"column:note;"`
	IsFavorite            bool           `json:"is_favorite" form:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived            bool           `json:"is_archived" form:"is_archived" gorm:"column:is_archived;"`
}

func (TemplateUpdate) TableName() string {
	return Template{}.TableName()
}

func (res *TemplateUpdate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.FakeCreatedBy = strings.TrimSpace(res.FakeCreatedBy)
	res.Note = strings.TrimSpace(res.Note)

	if res.FakeCollectionId != "" {
		res.FakeCollectionId = strings.TrimSpace(res.FakeCollectionId)
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
