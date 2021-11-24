package collectionmodel

import (
	"lift-tracker-api/common"
)

type CollectionCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" form:"name" binding:"required" gorm:"column:name;"`
	CreatedBy       int    `json:"-" gorm:"column:created_by;"`
	UserId          string `json:"created_by" form:"created_by" binding:"required" gorm:"-"`
	ParentId        string `json:"parent_id" form:"parent_id" gorm:"column:parent_id;default:null;"`
}

func (CollectionCreate) TableName() string {
	return Collection{}.TableName()
}

func (data *CollectionCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCollection)
}

func (res *CollectionCreate) Validate() error {
	return nil
}
