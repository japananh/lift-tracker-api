package templatebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

type UpdateTemplateStore interface {
	FindTemplateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*templatemodel.Template, error)
	Update(ctx context.Context, id int, data *templatemodel.TemplateUpdate) error
}

type updateTemplateBiz struct {
	store UpdateTemplateStore
}

func NewUpdateTemplateBiz(store UpdateTemplateStore) *updateTemplateBiz {
	return &updateTemplateBiz{store: store}
}

func (biz *updateTemplateBiz) UpdateTemplate(
	ctx context.Context,
	id int,
	data *templatemodel.TemplateUpdate,
) error {
	oldData, err := biz.store.FindTemplateByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(templatemodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(templatemodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
