package collectionbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

type ListCollectionStore interface {
	ListCollectionByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *collectionmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]collectionmodel.Collection, error)
}

type ListCollectionBiz struct {
	store ListCollectionStore
}

func NewListCollectionBiz(store ListCollectionStore) *ListCollectionBiz {
	return &ListCollectionBiz{store: store}
}

func (biz *ListCollectionBiz) ListCollection(
	ctx context.Context,
	filter *collectionmodel.Filter,
	paging *common.Paging,
) ([]collectionmodel.Collection, error) {
	result, err := biz.store.ListCollectionByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(collectionmodel.EntityName, err)
	}

	return result, nil
}
