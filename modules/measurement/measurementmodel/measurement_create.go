package measurementmodel

import (
	"lift-tracker-api/common"
	"strings"
)

type MeasurementCreate struct {
	common.SQLModel `json:",inline"`
	BodyPart        string `json:"body_part" form:"body_part" binding:"required" gorm:"column:body_part;type:enum('height', 'weight', 'bodyfat', 'neck', 'chest', 'left_biceps', 'right_biceps', 'left_forearms', 'right_forearms', 'waist', 'hips', 'glutes', 'left_thigh', 'right_thigh', 'left_calf', 'right_calf');"`
	FakeUserId      string `json:"user_id" form:"user_id" binding:"required" gorm:"-"`
	Unit            string `json:"unit" form:"unit" binding:"required" gorm:"column:unit;type:enum('kg', 'ibs', 'cm', 'in', '%');"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Value           int    `json:"value" form:"value" binding:"required" gorm:"column:value;"`
}

func (MeasurementCreate) TableName() string {
	return Measurement{}.TableName()
}

func (data *MeasurementCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeMeasurement)
}

func (res *MeasurementCreate) Validate() error {
	res.BodyPart = strings.TrimSpace(res.BodyPart)
	res.Unit = strings.TrimSpace(res.Unit)
	res.FakeUserId = strings.TrimSpace(res.FakeUserId)

	if res.FakeUserId != "" {
		res.FakeUserId = strings.TrimSpace(res.FakeUserId)

		// TODO: check for user not existed or disabled
		userId, err := common.FromBase58(strings.TrimSpace(res.FakeUserId))
		if err != nil {
			return err
		}

		res.UserId = int(userId.GetLocalID())
	}

	return nil
}
