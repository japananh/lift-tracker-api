package measurementbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

type GetMeasurementStore interface {
	FindMeasurementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*measurementmodel.Measurement, error)
}

type getMeasurementBiz struct {
	store GetMeasurementStore
}

func NewMeasurementBiz(store GetMeasurementStore) *getMeasurementBiz {
	return &getMeasurementBiz{store: store}
}

func (biz *getMeasurementBiz) GetMeasurement(
	ctx context.Context,
	id int,
) (*measurementmodel.Measurement, error) {
	data, err := biz.store.FindMeasurementByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(measurementmodel.EntityName, err)
	}

	return data, err
}
