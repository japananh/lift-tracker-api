package exercisestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

func (s *sqlStore) ListExerciseByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *exercisemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]exercisemodel.Exercise, error) {
	db := s.db
	db = db.Table(exercisemodel.Exercise{}.TableName()).
		Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.BodyParts != "" {
			db = db.Where("body_parts = ?", v.BodyParts)
		}

		if v.Mechanics != "" {
			db = db.Where("mechanics = ?", v.Mechanics)
		}

		if v.Force != "" {
			db = db.Where("force = ?", v.Force)
		}
	}

	// db.Count should come before db.Preload
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var result []exercisemodel.Exercise

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
