package exercisestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindExerciseByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*exercisemodel.Exercise, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var exercise exercisemodel.Exercise

	if err := db.
		Where(conditions).
		First(&exercise).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &exercise, nil
}
