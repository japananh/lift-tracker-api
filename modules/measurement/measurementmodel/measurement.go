package measurementmodel

import "lift-tracker-api/common"

const EntityName = "Measurement"

type Measurement struct {
	common.SQLModel `json:",inline"`
	BodyPart        string      `json:"body_part" gorm:"column:body_part;"`
	Unit            string      `json:"unit" gorm:"column:unit;"`
	UserId          int         `json:"-" gorm:"column:user_id;"`
	FakeUserId      *common.UID `json:"user_id" binding:"required" gorm:"-"`
	Value           int         `json:"value" gorm:"column:value;"`
}

func (Measurement) TableName() string {
	return "measurements"
}

func (data *Measurement) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeMeasurement)
	fakeUserId := common.NewUID(uint32(data.UserId), common.DbTypeMeasurement, 1)
	data.FakeUserId = &fakeUserId
}
