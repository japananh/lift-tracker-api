package collectionbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
	"lift-tracker-api/modules/user/usermodel"
)

type createCollectionStore interface {
	FindCollection(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*collectionmodel.Collection, error)
	CreateCollection(ctx context.Context, data *collectionmodel.CollectionCreate) error
}

type createCollectionBiz struct {
	store createCollectionStore
}

func NewCreateCollectionBiz(store createCollectionStore) *createCollectionBiz {
	return &createCollectionBiz{store: store}
}

func (biz *createCollectionBiz) CreateCollection(ctx context.Context, data *collectionmodel.CollectionCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	collection, err := biz.store.FindCollection(ctx, map[string]interface{}{"id": data.Id})

	if collection != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	if err == common.ErrRecordNotFound {
		if err := biz.store.CreateCollection(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(collectionmodel.EntityName, err)
		}
	} else {
		return common.ErrDB(err)
	}

	return nil
}
