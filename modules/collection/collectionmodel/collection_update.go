package collectionmodel

import (
	"lift-tracker-api/common"
	"strings"
)

type CollectionUpdate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" form:"name" gorm:"column:name;"`
	IsFavorite      bool   `json:"is_favorite" form:"is_favorite" gorm:"column:is_favorite;"`
	IsArchived      bool   `json:"is_archived" form:"is_archived" gorm:"column:is_archived;"`
	ParentId        int    `json:"-" gorm:"column:parent_id;"`
	FakeParentId    string `json:"parent_id" form:"parent_id" gorm:"-"`
}

func (CollectionUpdate) TableName() string {
	return Collection{}.TableName()
}

func (res *CollectionUpdate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.FakeParentId = strings.TrimSpace(res.FakeParentId)

	if res.FakeParentId != "" {
		parentId, err := common.FromBase58(strings.TrimSpace(res.FakeParentId))
		if err != nil {
			return err
		}

		res.ParentId = int(parentId.GetLocalID())
	}

	return nil
}
