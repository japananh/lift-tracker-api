package measurementbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

type DeleteMeasurementStore interface {
	FindMeasurementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*measurementmodel.Measurement, error)
	Delete(ctx context.Context, id int) error
}

type deleteMeasurementBiz struct {
	store DeleteMeasurementStore
}

func NewDeleteMeasurementBiz(store DeleteMeasurementStore) *deleteMeasurementBiz {
	return &deleteMeasurementBiz{store: store}
}

func (biz *deleteMeasurementBiz) DeleteMeasurement(
	ctx context.Context,
	id int,
) error {
	data, err := biz.store.FindMeasurementByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(measurementmodel.EntityName, nil)
	}

	if data.Status == 0 {
		return common.ErrEntityDeleted(measurementmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
