package collectionmodel

import (
	"lift-tracker-api/common"
	"strings"
)

type CollectionCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" form:"name" binding:"required" gorm:"column:name;"`
	CreatedBy       int    `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy   string `json:"created_by" form:"created_by" binding:"required" gorm:"-"`
	ParentId        int    `json:"-" gorm:"column:parent_id;default:null;"`
	FakeParentId    string `json:"parent_id" form:"parent_id" gorm:"-"`
}

func (CollectionCreate) TableName() string {
	return Collection{}.TableName()
}

func (data *CollectionCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCollection)
}

func (res *CollectionCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if res.FakeCreatedBy != "" {
		res.FakeCreatedBy = strings.TrimSpace(res.FakeCreatedBy)

		// TODO: check for user not existed or disabled
		createdBy, err := common.FromBase58(strings.TrimSpace(res.FakeCreatedBy))
		if err != nil {
			return err
		}

		res.CreatedBy = int(createdBy.GetLocalID())
	}

	if res.FakeParentId != "" {
		res.FakeParentId = strings.TrimSpace(res.FakeParentId)

		parentId, err := common.FromBase58(strings.TrimSpace(res.FakeParentId))
		if err != nil {
			return err
		}

		res.ParentId = int(parentId.GetLocalID())
	}

	return nil
}
