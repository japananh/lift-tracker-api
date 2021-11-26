package collectionbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

type UpdateCollectionStore interface {
	FindCollectionByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*collectionmodel.Collection, error)
	Update(ctx context.Context, id int, data *collectionmodel.CollectionUpdate) error
}

type updateCollectionBiz struct {
	store UpdateCollectionStore
}

func NewUpdateCollectionBiz(store UpdateCollectionStore) *updateCollectionBiz {
	return &updateCollectionBiz{store: store}
}

func (biz *updateCollectionBiz) UpdateCollection(
	ctx context.Context,
	id int,
	data *collectionmodel.CollectionUpdate,
) error {
	oldData, err := biz.store.FindCollectionByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(collectionmodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(collectionmodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
