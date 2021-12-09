package templatemodel

import "lift-tracker-api/common"

type Filter struct {
	CreatedBy        int         `json:"-"  gorm:"column:created_by;"`
	FakeCreatedBy    *common.UID `json:"created_by" gorm:"-"`
	CollectionId     int         `json:"-" gorm:"column:collection_id;"`
	FakeCollectionId *common.UID `json:"collection_id"  gorm:"-"`
}
