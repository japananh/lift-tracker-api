package exercisemodel

import (
	"lift-tracker-api/common"
	"strings"
)

type ExerciseUpdate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" form:"name" gorm:"column:name;"`
	CreatedBy       int           `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy   string        `json:"created_by" form:"created_by" gorm:"-"`
	Category        string        `json:"category" form:"category" gorm:"column:category;type:enum('barbell', 'dumbbell', 'machine/other', 'weighted bodyweight', 'assisted body', 'reps only', 'cardio exercise', 'duration');"`
	BodyParts       string        `json:"body_parts" form:"body_parts" gorm:"column:body_parts;"`
	Mechanics       string        `json:"mechanics" form:"mechanics" gorm:"column:mechanics;type:enum('isolation', 'compound');"`
	Force           string        `json:"force" form:"force" gorm:"column:force;type:enum('push', 'pull');"`
	Instructions    string        `json:"instructions" form:"instructions" gorm:"column:instructions;"`
	RestTime        int           `json:"rest_time" form:"rest_time" gorm:"column:rest_time;"`
	Image           *common.Image `json:"image,omitempty" gorm:"column:image;type:json"`
}

func (ExerciseUpdate) TableName() string {
	return Exercise{}.TableName()
}

func (res *ExerciseUpdate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.FakeCreatedBy = strings.TrimSpace(res.FakeCreatedBy)

	if res.FakeCreatedBy != "" {
		createdBy, err := common.FromBase58(strings.TrimSpace(res.FakeCreatedBy))
		if err != nil {
			return err
		}

		res.CreatedBy = int(createdBy.GetLocalID())
	}

	return nil
}
