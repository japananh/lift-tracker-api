package usermodel

import (
	"errors"
	"lift-tracker-api/common"
)

const EntityName = "User"

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)

func ErrPasswordInvalid(msg string) *common.AppError {
	return common.NewCustomError(
		errors.New(msg),
		msg,
		"ErrPasswordInvalid",
	)
}

type User struct {
	common.SQLModel `json:",inline"`
	Status          int           `json:"status" gorm:"column:status;default:1;"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Role            string        `json:"role" gorm:"column:role;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (data *User) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeUser)
}
