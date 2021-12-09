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

func ListTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter templatemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		paging.Fulfill()

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := templatestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := templatebiz.NewListTemplateBiz(store)

		result, err := biz.ListTemplate(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
