package exercisebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

type DeleteExerciseStore interface {
	FindExerciseByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*exercisemodel.Exercise, error)
	Delete(ctx context.Context, id int) error
}

type deleteExerciseBiz struct {
	store DeleteExerciseStore
}

func NewDeleteExerciseBiz(store DeleteExerciseStore) *deleteExerciseBiz {
	return &deleteExerciseBiz{store: store}
}

func (biz *deleteExerciseBiz) DeleteExercise(
	ctx context.Context,
	id int,
) error {
	data, err := biz.store.FindExerciseByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(exercisemodel.EntityName, nil)
	}

	if data.Status == 0 {
		return common.ErrEntityDeleted(exercisemodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
