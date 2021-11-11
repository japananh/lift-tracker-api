package userbiz

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/component/tokenprovider"
	"lift-tracker-api/modules/user/usermodel"
)

type LoginStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	loginStore    LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tokenConfig   *tokenprovider.TokenConfig
}

func NewLoginBiz(
	loginStore LoginStore,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	tokenConfig *tokenprovider.TokenConfig,
) *loginBiz {
	return &loginBiz{
		loginStore:    loginStore,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tokenConfig:   tokenConfig,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	user, err := biz.loginStore.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	hashedPassword := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != hashedPassword {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tokenConfig.AccessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.tokenConfig.RefreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
