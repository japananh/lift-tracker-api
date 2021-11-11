package usermodel

import (
	"lift-tracker-api/common"
	"strings"
	"unicode"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Status          int           `json:"status" gorm:"column:status;default:1;"`
	Email           string        `json:"email" form:"email" binding:"required" gorm:"column:email;"`
	Password        string        `json:"password" form:"password" binding:"required" gorm:"column:password;"`
	FirstName       string        `json:"first_name" form:"first_name" gorm:"first_name;"`
	LastName        string        `json:"last_name" form:"last_name" gorm:"last_name;"`
	Role            string        `json:"role" form:"role" gorm:"column:role;type:enum('user', 'admin');default:'user'"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" form:"avatar" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeUser)
}

func (res *UserCreate) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)

	if errMsg := verifyPassword(res.Password); errMsg != "" {
		return ErrPasswordInvalid(errMsg)
	}

	return nil
}

func verifyPassword(s string) string {
	hasNumber, hasLetter, hasSpecial, hasInvalidCharacter := false, false, false, false
	letterCount := 0

	for _, c := range s {
		letterCount++
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLetter(c):
			hasLetter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		default:
			hasInvalidCharacter = true
		}
	}

	if hasInvalidCharacter {
		return "password has invalid characters"
	}

	if letterCount < 8 {
		return "password must have at least 8 characters"
	}

	if !hasNumber {
		return "password must have at least 1 number"
	}

	if !hasLetter {
		return "password must have at least 1 letter"
	}

	if !hasSpecial {
		return "password must have at least 1 speicial character"
	}

	return ""
}
