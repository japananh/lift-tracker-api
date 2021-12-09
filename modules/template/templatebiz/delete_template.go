package templatebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

type DeleteTemplateStore interface {
	FindTemplateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*templatemodel.Template, error)
	Delete(ctx context.Context, id int) error
}

type deleteTemplateBiz struct {
	store DeleteTemplateStore
}

func NewDeleteTemplateBiz(store DeleteTemplateStore) *deleteTemplateBiz {
	return &deleteTemplateBiz{store: store}
}

func (biz *deleteTemplateBiz) DeleteTemplate(
	ctx context.Context,
	id int,
) error {
	data, err := biz.store.FindTemplateByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(templatemodel.EntityName, nil)
	}

	if data.Status == 0 {
		return common.ErrEntityDeleted(templatemodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
