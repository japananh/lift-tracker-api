package component

import (
	"lift-tracker-api/component/tokenprovider"
	"lift-tracker-api/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetTokenConfig() *tokenprovider.TokenConfig
}

type appCtx struct {
	db          *gorm.DB
	upProvider  uploadprovider.UploadProvider
	secretKey   string
	tokenConfig *tokenprovider.TokenConfig
}

func NewAppContext(
	db *gorm.DB,
	upProvider uploadprovider.UploadProvider,
	secretKey string,
	tokenConfig *tokenprovider.TokenConfig,
) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey, tokenConfig: tokenConfig}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetTokenConfig() *tokenprovider.TokenConfig {
	return ctx.tokenConfig
}
