package templatemodel

import (
	"lift-tracker-api/common"

	"gorm.io/datatypes"
)

const EntityName = "Template"

type Template struct {
	common.SQLModel  `json:",inline"`
	Name             string         `json:"name" gorm:"column:name;"`
	CreatedBy        int            `json:"-"  gorm:"column:created_by;"`
	FakeCreatedBy    *common.UID    `json:"created_by" gorm:"-"`
	Note             string         `json:"note" gorm:"column:note;"`
	Detail           datatypes.JSON `json:"detail" form:"detail" binding:"required" gorm:"column:detail;"`
	CollectionId     int            `json:"-" gorm:"column:collection_id;"`
	FakeCollectionId *common.UID    `json:"collection_id"  gorm:"-"`
	IsFavorite       bool           `json:"is_favorite" form:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived       bool           `json:"is_archived" form:"is_archived" gorm:"column:is_archived;"`
}

func (Template) TableName() string {
	return "templates"
}

func (data *Template) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeTemplate)
	fakeCreatedBy := common.NewUID(uint32(data.CreatedBy), common.DbTypeTemplate, 1)
	data.FakeCreatedBy = &fakeCreatedBy

	fakeCollectionId := common.NewUID(uint32(data.CollectionId), common.DbTypeTemplate, 1)
	data.FakeCollectionId = &fakeCollectionId
}
