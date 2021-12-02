package exercisebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

type GetExerciseStore interface {
	FindExerciseByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*exercisemodel.Exercise, error)
}

type getExerciseBiz struct {
	store GetExerciseStore
}

func NewExerciseBiz(store GetExerciseStore) *getExerciseBiz {
	return &getExerciseBiz{store: store}
}

func (biz *getExerciseBiz) GetExercise(
	ctx context.Context,
	id int,
) (*exercisemodel.Exercise, error) {
	data, err := biz.store.FindExerciseByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(exercisemodel.EntityName, err)
	}

	return data, err
}
