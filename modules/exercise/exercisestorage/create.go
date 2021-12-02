package exercisestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

func (s *sqlStore) CreateExercise(ctx context.Context, data *exercisemodel.ExerciseCreate) error {
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
