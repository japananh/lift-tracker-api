package common

const (
	DbTypeUser        = 1
	DbTypeWorkout     = 2
	DbTypeTemplate    = 3
	DbTypeRecord      = 4
	DbTypeCollection  = 5
	DbTypeExercise    = 6
	DbTypeSetting     = 7
	DbTypeMeasurement = 8
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
