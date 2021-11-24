package collectionstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

func (s *sqlStore) CreateCollection(ctx context.Context, data *collectionmodel.CollectionCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
