package ginexercise

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/exercise/exercisebiz"
	"lift-tracker-api/modules/exercise/exercisemodel"
	"lift-tracker-api/modules/exercise/exercisestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListExercise(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter exercisemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		paging.Fulfill()

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := exercisestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := exercisebiz.NewListExerciseBiz(store)

		result, err := biz.ListExercise(c.Request.Context(), &filter, &paging)
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
