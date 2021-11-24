package collectionstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindCollection(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*collectionmodel.Collection, error) {
	db := s.db.Table(collectionmodel.Collection{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var collection collectionmodel.Collection

	if err := db.Where(conditions).First(&collection).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &collection, nil
}
