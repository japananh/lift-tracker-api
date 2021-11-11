package ginuser

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// type conversion to Requester, so no need to import user
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
