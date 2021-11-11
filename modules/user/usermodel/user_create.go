package usermodel

import (
	"lift-tracker-api/common"
	"strings"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" form:"email" gorm:"column:email;"`
	Password        string        `json:"password" form:"password" gorm:"column:password;"`
	FirstName       string        `json:"first_name" form:"first_name" gorm:"first_name;"`
	LastName        string        `json:"last_name" form:"last_name" gorm:"last_name;"`
	Role            string        `json:"role" form:"role" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" form:"avatar" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (res *UserCreate) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)

	if len(res.Email) == 0 || len(res.Password) == 0 {
		return ErrEmailOrPasswordInvalid
	}

	return nil
}

func (data *UserCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeUser)
}
