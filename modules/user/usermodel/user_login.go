package usermodel

import (
	"lift-tracker-api/component/tokenprovider"
	"strings"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

func (res *UserLogin) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)

	if len(res.Email) == 0 || len(res.Password) == 0 {
		return ErrEmailOrPasswordInvalid
	}

	return nil
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}
