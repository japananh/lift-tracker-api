package measurementbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/measurement/measurementmodel"
)

type ListMeasurementStore interface {
	ListMeasurementByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *measurementmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]measurementmodel.Measurement, error)
}

type ListMeasurementBiz struct {
	store ListMeasurementStore
}

func NewListMeasurementBiz(store ListMeasurementStore) *ListMeasurementBiz {
	return &ListMeasurementBiz{store: store}
}

func (biz *ListMeasurementBiz) ListMeasurement(
	ctx context.Context,
	filter *measurementmodel.Filter,
	paging *common.Paging,
) ([]measurementmodel.Measurement, error) {
	result, err := biz.store.ListMeasurementByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(measurementmodel.EntityName, err)
	}

	return result, nil
}
