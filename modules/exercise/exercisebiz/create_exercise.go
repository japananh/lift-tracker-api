package exercisebiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/exercise/exercisemodel"
)

type createExerciseStore interface {
	FindExerciseByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*exercisemodel.Exercise, error)
	CreateExercise(ctx context.Context, data *exercisemodel.ExerciseCreate) error
}

type createExerciseBiz struct {
	store createExerciseStore
}

func NewCreateExerciseBiz(store createExerciseStore) *createExerciseBiz {
	return &createExerciseBiz{store: store}
}

func (biz *createExerciseBiz) CreateExercise(ctx context.Context, data *exercisemodel.ExerciseCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// TODO: check for invalid createdBy
	duplicatedExercise, err := biz.store.FindExerciseByCondition(ctx, map[string]interface{}{"id": data.Id})
	if duplicatedExercise != nil {
		return common.ErrEntityExisted(exercisemodel.EntityName, err)
	}

	if err == common.ErrRecordNotFound {
		if err := biz.store.CreateExercise(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(exercisemodel.EntityName, err)
		}
	} else {
		return common.ErrDB(err)
	}

	return nil
}
