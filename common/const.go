package common

const (
	DbTypeUser             = 1
	DbTypeWorkout          = 2
	DbTypeTemplate         = 3
	DbTypeRecord           = 4
	DbTypeDirectory        = 5
	DbTypeExercise         = 6
	DbTypeSetting          = 7
	DbTypeMeasurement      = 8
	DbTypeBodyPart         = 9
	DbTypeExerciseBodyPart = 10
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
