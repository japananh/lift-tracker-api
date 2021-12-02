package ginmeasurement

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/measurement/measurementbiz"
	"lift-tracker-api/modules/measurement/measurementstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteMeasurement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := measurementstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := measurementbiz.NewDeleteMeasurementBiz(store)

		if err := biz.DeleteMeasurement(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
