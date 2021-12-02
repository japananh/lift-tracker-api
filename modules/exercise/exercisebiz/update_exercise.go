package exercisebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

type UpdateExerciseStore interface {
	FindExerciseByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*exercisemodel.Exercise, error)
	Update(ctx context.Context, id int, data *exercisemodel.ExerciseUpdate) error
}

type updateExerciseBiz struct {
	store UpdateExerciseStore
}

func NewUpdateExerciseBiz(store UpdateExerciseStore) *updateExerciseBiz {
	return &updateExerciseBiz{store: store}
}

func (biz *updateExerciseBiz) UpdateCollection(
	ctx context.Context,
	id int,
	data *exercisemodel.ExerciseUpdate,
) error {
	oldData, err := biz.store.FindExerciseByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(exercisemodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(exercisemodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
