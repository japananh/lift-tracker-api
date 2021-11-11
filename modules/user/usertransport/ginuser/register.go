package ginuser

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/component/hasher"
	"lift-tracker-api/modules/user/userbiz"
	"lift-tracker-api/modules/user/usermodel"
	"lift-tracker-api/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
