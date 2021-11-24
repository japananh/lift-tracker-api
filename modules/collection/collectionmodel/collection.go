package collectionmodel

import "lift-tracker-api/common"

const EntityName = "Collection"

type Collection struct {
	common.SQLModel `json:",inline"`
	ParentId        string `json:"parent_id" gorm:"column:parent_id;"`
	Name            string `json:"name" gorm:"column:name;"`
	IsFavorite      string `json:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived      string `json:"is_archived" gorm:"column:is_archived;"`
	CreatedBy       string `json:"created_by" gorm:"column:created_by;"`
}

func (Collection) TableName() string {
	return "collections"
}

func (data *Collection) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCollection)
}
