package collectionstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

func (s *sqlStore) Delete(ctx context.Context, ids []int) error {
	db := s.db

	if err := db.
		Table(collectionmodel.Collection{}.TableName()).
		Delete(&collectionmodel.Collection{}, "id in ?", ids).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
