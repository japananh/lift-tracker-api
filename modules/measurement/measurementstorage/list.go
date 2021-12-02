package measurementstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

func (s *sqlStore) ListMeasurementByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *measurementmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]measurementmodel.Measurement, error) {
	db := s.db
	db = db.Table(measurementmodel.Measurement{}.TableName()).
		Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
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

	var result []measurementmodel.Measurement

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
