package templatebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

type ListTemplateStore interface {
	ListTemplateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *templatemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]templatemodel.Template, error)
}

type ListTemplateBiz struct {
	store ListTemplateStore
}

func NewListTemplateBiz(store ListTemplateStore) *ListTemplateBiz {
	return &ListTemplateBiz{store: store}
}

func (biz *ListTemplateBiz) ListTemplate(
	ctx context.Context,
	filter *templatemodel.Filter,
	paging *common.Paging,
) ([]templatemodel.Template, error) {
	result, err := biz.store.ListTemplateByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(templatemodel.EntityName, err)
	}

	return result, nil
}
