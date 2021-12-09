package templatebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

type createTemplateStore interface {
	FindTemplateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*templatemodel.Template, error)
	CreateTemplate(ctx context.Context, data *templatemodel.TemplateCreate) error
}

type createTemplateBiz struct {
	store createTemplateStore
}

func NewCreateTemplateBiz(store createTemplateStore) *createTemplateBiz {
	return &createTemplateBiz{store: store}
}

func (biz *createTemplateBiz) CreateTemplate(ctx context.Context, data *templatemodel.TemplateCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// TODO: check for invalid createdBy
	duplicatedTemplate, err := biz.store.FindTemplateByCondition(ctx, map[string]interface{}{"id": data.Id})
	if duplicatedTemplate != nil {
		return common.ErrEntityExisted(templatemodel.EntityName, err)
	}

	if err == common.ErrRecordNotFound {
		if err := biz.store.CreateTemplate(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(templatemodel.EntityName, err)
		}
	} else {
		return common.ErrDB(err)
	}

	return nil
}
