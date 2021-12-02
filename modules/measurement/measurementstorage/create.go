package measurementstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

func (s *sqlStore) CreateMeasurement(ctx context.Context, data *measurementmodel.MeasurementCreate) error {
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
