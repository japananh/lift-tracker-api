package measurementmodel

import (
	"lift-tracker-api/common"
	"strings"
)

type MeasurementUpdate struct {
	common.SQLModel `json:",inline"`
	BodyPart        string `json:"body_part" form:"body_part" gorm:"column:body_part;type:enum('height', 'weight', 'bodyfat', 'neck', 'chest', 'left_biceps', 'right_biceps', 'left_forearms', 'right_forearms', 'waist', 'hips', 'glutes', 'left_thigh', 'right_thigh', 'left_calf', 'right_calf');"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	FakeUserId      string `json:"user_id" form:"user_id" gorm:"-"`
	Value           int    `json:"value" form:"value" gorm:"column:value;"`
	Unit            string `json:"unit" form:"unit" gorm:"column:unit;type:enum('kg', 'ibs', 'cm', 'in');"`
}

func (MeasurementUpdate) TableName() string {
	return Measurement{}.TableName()
}

func (res *MeasurementUpdate) Validate() error {
	res.BodyPart = strings.TrimSpace(res.BodyPart)
	res.Unit = strings.TrimSpace(res.Unit)
	res.FakeUserId = strings.TrimSpace(res.FakeUserId)

	if res.FakeUserId != "" {
		userId, err := common.FromBase58(strings.TrimSpace(res.FakeUserId))
		if err != nil {
			return err
		}

		res.UserId = int(userId.GetLocalID())
	}

	return nil
}
