package gintemplate

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/template/templatebiz"
	"lift-tracker-api/modules/template/templatestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := templatestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := templatebiz.NewTemplateBiz(store)

		data, err := biz.GetTemplate(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
