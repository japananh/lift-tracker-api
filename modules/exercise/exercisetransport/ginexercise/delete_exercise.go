package ginexercise

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/exercise/exercisebiz"
	"lift-tracker-api/modules/exercise/exercisestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteExercise(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := exercisestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := exercisebiz.NewDeleteExerciseBiz(store)

		if err := biz.DeleteExercise(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
