package collectionbiz

import (
	"context"
	"lift-tracker-api/modules/collection/collectionmodel"
)

type DeleteRestaurantStore interface {
	FindCollectionByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*collectionmodel.Collection, error)
	Delete(ctx context.Context, ids []int) error
}

type deleteCollectionBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteCollectionBiz(store DeleteRestaurantStore) *deleteCollectionBiz {
	return &deleteCollectionBiz{store: store}
}

func (biz *deleteCollectionBiz) DeleteCollection(
	ctx context.Context,
	ids []int,
) ([]int, error) {
	var invalidIds []int

	for _, id := range ids {
		_, err := biz.store.FindCollectionByCondition(ctx, map[string]interface{}{"id": id})
		if err != nil {
			// return common.ErrCannotGetEntity(collectionmodel.EntityName, nil)
			invalidIds = append(invalidIds, id)
		}
	}

	if err := biz.store.Delete(ctx, ids); err != nil {
		return nil, err
	}

	return invalidIds, nil
}
