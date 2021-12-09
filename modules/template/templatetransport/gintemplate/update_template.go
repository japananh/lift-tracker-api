package gintemplate

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/template/templatebiz"
	"lift-tracker-api/modules/template/templatemodel"
	"lift-tracker-api/modules/template/templatestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data templatemodel.TemplateUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := templatestorage.NewSQLStore(db)
		biz := templatebiz.NewUpdateTemplateBiz(store)

		if err := biz.UpdateTemplate(
			c.Request.Context(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
