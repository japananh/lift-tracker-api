package exercisebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

type ListExerciseStore interface {
	ListExerciseByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *exercisemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]exercisemodel.Exercise, error)
}

type ListExerciseBiz struct {
	store ListExerciseStore
}

func NewListExerciseBiz(store ListExerciseStore) *ListExerciseBiz {
	return &ListExerciseBiz{store: store}
}

func (biz *ListExerciseBiz) ListExercise(
	ctx context.Context,
	filter *exercisemodel.Filter,
	paging *common.Paging,
) ([]exercisemodel.Exercise, error) {
	result, err := biz.store.ListExerciseByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(exercisemodel.EntityName, err)
	}

	return result, nil
}
