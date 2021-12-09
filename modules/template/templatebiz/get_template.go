package templatebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

type GetTemplateStore interface {
	FindTemplateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*templatemodel.Template, error)
}

type getTemplateBiz struct {
	store GetTemplateStore
}

func NewTemplateBiz(store GetTemplateStore) *getTemplateBiz {
	return &getTemplateBiz{store: store}
}

func (biz *getTemplateBiz) GetTemplate(
	ctx context.Context,
	id int,
) (*templatemodel.Template, error) {
	data, err := biz.store.FindTemplateByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(templatemodel.EntityName, err)
	}

	return data, err
}
