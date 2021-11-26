package collectionmodel

import "lift-tracker-api/common"

const EntityName = "Collection"

type Collection struct {
	common.SQLModel `json:",inline"`
	Name            string      `json:"name" gorm:"column:name;"`
	IsFavorite      bool        `json:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived      bool        `json:"is_archived" gorm:"column:is_archived;"`
	CreatedBy       int         `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy   *common.UID `json:"created_by" binding:"required" gorm:"-"`
	ParentId        int         `json:"-" gorm:"column:parent_id;"`
	FakeParentId    *common.UID `json:"parent_id" gorm:"-"`
}

func (Collection) TableName() string {
	return "collections"
}

func (data *Collection) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCollection)
	fakeCreatedBy := common.NewUID(uint32(data.CreatedBy), common.DbTypeCollection, 1)
	data.FakeCreatedBy = &fakeCreatedBy
	parentId := common.NewUID(uint32(data.ParentId), common.DbTypeCollection, 1)
	data.FakeParentId = &parentId
}
