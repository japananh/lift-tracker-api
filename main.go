package main

import (
	"lift-tracker-api/component"
	"lift-tracker-api/component/tokenprovider"
	"lift-tracker-api/component/uploadprovider"
	"lift-tracker-api/middleware"
	"lift-tracker-api/modules/user/usertransport/ginuser"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION_STR")
	s3bucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3ApiKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := os.Getenv("S3_DOMAIN")
	secretKey := os.Getenv("SYSTEM_KEY")
	atExpiryStr := os.Getenv("ACCESS_TOKEN_EXPIRY")
	rtExpiryStr := os.Getenv("REFRESH_TOKEN_EXPIRY")

	if dsn == "" || s3bucketName == "" || s3Region == "" || s3ApiKey == "" ||
		s3Secret == "" || atExpiryStr == "" || rtExpiryStr == "" {
		log.Fatalln("env not found")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	s3Provider := uploadprovider.NewS3Provider(s3bucketName, s3Region, s3ApiKey, s3Secret, s3Domain)
	tokenConfig, err := tokenprovider.NewTokenConfig(atExpiryStr, rtExpiryStr)
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider, secretKey, tokenConfig); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB,
	upProvider uploadprovider.UploadProvider,
	secretKey string,
	tokenConfig *tokenprovider.TokenConfig,
) error {
	r := gin.Default()

	appCtx := component.NewAppContext(db, upProvider, secretKey, tokenConfig)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/login", ginuser.Login(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/register/verify", ginuser.VerifyEmail(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	return r.Run()
}
