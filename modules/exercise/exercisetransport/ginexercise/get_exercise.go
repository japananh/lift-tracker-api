package ginexercise

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/exercise/exercisebiz"
	"lift-tracker-api/modules/exercise/exercisestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExercise(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := exercisestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := exercisebiz.NewExerciseBiz(store)

		data, err := biz.GetExercise(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
