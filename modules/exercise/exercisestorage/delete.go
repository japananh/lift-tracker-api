package exercisestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.
		Table(exercisemodel.Exercise{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
