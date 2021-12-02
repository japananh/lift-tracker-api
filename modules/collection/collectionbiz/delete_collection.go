package collectionbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

type DeleteRestaurantStore interface {
	FindCollectionByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*collectionmodel.Collection, error)
	Delete(ctx context.Context, id int) error
}

type deleteCollectionBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteCollectionBiz(store DeleteRestaurantStore) *deleteCollectionBiz {
	return &deleteCollectionBiz{store: store}
}

func (biz *deleteCollectionBiz) DeleteCollection(
	ctx context.Context,
	id int,
) error {
	data, err := biz.store.FindCollectionByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(collectionmodel.EntityName, nil)
	}

	if data.Status == 0 {
		return common.ErrEntityDeleted(collectionmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
