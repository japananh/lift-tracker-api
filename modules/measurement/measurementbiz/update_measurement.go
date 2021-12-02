package measurementbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

type UpdateMeasurementStore interface {
	FindMeasurementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*measurementmodel.Measurement, error)
	Update(ctx context.Context, id int, data *measurementmodel.MeasurementUpdate) error
}

type updateMeasurementBiz struct {
	store UpdateMeasurementStore
}

func NewUpdateMeasurementBiz(store UpdateMeasurementStore) *updateMeasurementBiz {
	return &updateMeasurementBiz{store: store}
}

func (biz *updateMeasurementBiz) UpdateMeasurement(
	ctx context.Context,
	id int,
	data *measurementmodel.MeasurementUpdate,
) error {
	oldData, err := biz.store.FindMeasurementByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(measurementmodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(measurementmodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
