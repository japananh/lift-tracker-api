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

func CreateExercise(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data exercisemodel.ExerciseCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := exercisestorage.NewSQLStore(db)
		biz := exercisebiz.NewCreateExerciseBiz(store)

		if err := biz.CreateExercise(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
