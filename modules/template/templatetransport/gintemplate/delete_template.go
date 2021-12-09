package gintemplate

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/template/templatebiz"
	"lift-tracker-api/modules/template/templatestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := templatestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := templatebiz.NewDeleteTemplateBiz(store)

		if err := biz.DeleteTemplate(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
