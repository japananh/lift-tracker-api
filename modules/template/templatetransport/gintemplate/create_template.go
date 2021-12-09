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

func CreateTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data templatemodel.TemplateCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := templatestorage.NewSQLStore(db)
		biz := templatebiz.NewCreateTemplateBiz(store)

		if err := biz.CreateTemplate(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
