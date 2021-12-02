package measurementbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
	"lift-tracker-api/modules/user/usermodel"
)

type createMeasurementStore interface {
	FindMeasurementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*measurementmodel.Measurement, error)
	CreateMeasurement(ctx context.Context, data *measurementmodel.MeasurementCreate) error
}

type createMeasurementBiz struct {
	store createMeasurementStore
}

func NewCreateMeasurementBiz(store createMeasurementStore) *createMeasurementBiz {
	return &createMeasurementBiz{store: store}
}

func (biz *createMeasurementBiz) CreateMeasurement(ctx context.Context, data *measurementmodel.MeasurementCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	duplicatedMeasurement, err := biz.store.FindMeasurementByCondition(ctx, map[string]interface{}{"id": data.Id})
	if duplicatedMeasurement != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	if err == common.ErrRecordNotFound {
		if err := biz.store.CreateMeasurement(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(measurementmodel.EntityName, err)
		}
	} else {
		return common.ErrDB(err)
	}

	return nil
}
