package measurementstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindMeasurementByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*measurementmodel.Measurement, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var measurement measurementmodel.Measurement

	if err := db.
		Where(conditions).
		First(&measurement).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &measurement, nil
}
