package exercisemodel

import (
	"lift-tracker-api/common"
	"strings"
)

type ExerciseCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" form:"name" binding:"required" gorm:"column:name;"`
	CreatedBy       int           `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy   string        `json:"created_by" form:"created_by" binding:"required" gorm:"-"`
	Category        string        `json:"category" form:"category" binding:"required" gorm:"column:category;type:enum('barbell', 'dumbbell', 'machine/other', 'weighted bodyweight', 'assisted body', 'reps only', 'cardio exercise', 'duration');"`
	BodyParts       string        `json:"body_parts" form:"body_parts" binding:"required" gorm:"column:body_parts;"`
	Mechanics       string        `json:"mechanics" form:"mechanics" gorm:"column:mechanics;type:enum('isolation', 'compound');"`
	Force           string        `json:"force" form:"force" gorm:"column:force;type:enum('push', 'pull');"`
	Instructions    string        `json:"instructions" form:"instructions" gorm:"column:instructions;"`
	RestTime        int           `json:"rest_time" form:"rest_time" gorm:"column:rest_time;"`
	Image           *common.Image `json:"image,omitempty" gorm:"column:image;type:json"`
}

func (ExerciseCreate) TableName() string {
	return Exercise{}.TableName()
}

func (data *ExerciseCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeExercise)
}

func (res *ExerciseCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.Category = strings.TrimSpace(res.Category)

	// TODO: Validate body_parts
	// type:enum('arms', 'shoulders', 'glutes', 'legs', 'chest', 'core', 'back', 'full_body', 'cardio', 'other')
	res.BodyParts = strings.TrimSpace(res.BodyParts)

	if res.Mechanics != "" {
		res.Mechanics = strings.TrimSpace(res.Mechanics)
	}

	if res.Force != "" {
		res.Force = strings.TrimSpace(res.Force)
	}

	if res.Instructions != "" {
		res.Instructions = strings.TrimSpace(res.Instructions)
	}

	if res.FakeCreatedBy != "" {
		res.FakeCreatedBy = strings.TrimSpace(res.FakeCreatedBy)

		createdBy, err := common.FromBase58(strings.TrimSpace(res.FakeCreatedBy))
		if err != nil {
			return err
		}

		res.CreatedBy = int(createdBy.GetLocalID())
	}
	return nil
}
