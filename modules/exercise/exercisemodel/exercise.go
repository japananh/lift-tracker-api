package exercisemodel

import "lift-tracker-api/common"

const EntityName = "Exercise"

// TODO: handle store video
type Exercise struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	CreatedBy       int           `json:"-" gorm:"column:created_by;"`
	FakeCreatedBy   *common.UID   `json:"created_by" binding:"required" gorm:"-"`
	Category        string        `json:"category" gorm:"column:category;"`
	BodyParts       string        `json:"body_parts" gorm:"column:body_parts;"`
	Mechanics       string        `json:"mechanics" gorm:"column:mechanics;"`
	Force           string        `json:"force" gorm:"column:force;"`
	RestTime        int           `json:"rest_time" gorm:"column:rest_time;"`
	Instructions    string        `json:"instructions" gorm:"column:instructions;"`
	Image           *common.Image `json:"image,omitempty" gorm:"column:image;type:json"`
}

func (Exercise) TableName() string {
	return "exercises"
}

func (data *Exercise) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeExercise)
	fakeCreatedBy := common.NewUID(uint32(data.CreatedBy), common.DbTypeExercise, 1)
	data.FakeCreatedBy = &fakeCreatedBy
}
