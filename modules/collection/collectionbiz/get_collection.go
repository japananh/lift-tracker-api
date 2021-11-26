package collectionbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

type GetCollectionStore interface {
	FindCollectionByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*collectionmodel.Collection, error)
}

type getCollectionBiz struct {
	store GetCollectionStore
}

func NewCollectionBiz(store GetCollectionStore) *getCollectionBiz {
	return &getCollectionBiz{store: store}
}

func (biz *getCollectionBiz) GetCollection(
	ctx context.Context,
	id int,
) (*collectionmodel.Collection, error) {
	data, err := biz.store.FindCollectionByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(collectionmodel.EntityName, err)
	}

	return data, err
}
