package ginuser

import (
	"errors"
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/user/usermodel"
	"lift-tracker-api/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyEmail(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")

		if len(email) == 0 {
			panic(common.ErrInvalidRequest(errors.New("invalid email")))
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"email": email})

		if user != nil {
			panic(common.ErrEntityExisted(usermodel.EntityName, err))
		}

		c.JSON(http.StatusOK, "ok")
	}
}
