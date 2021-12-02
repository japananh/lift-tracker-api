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

func UpdateExercise(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data exercisemodel.ExerciseUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := exercisestorage.NewSQLStore(db)
		biz := exercisebiz.NewUpdateExerciseBiz(store)

		if err := biz.UpdateCollection(
			c.Request.Context(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
