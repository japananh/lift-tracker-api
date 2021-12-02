package ginmeasurement

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/measurement/measurementbiz"
	"lift-tracker-api/modules/measurement/measurementstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMeasurement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := measurementstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := measurementbiz.NewMeasurementBiz(store)

		data, err := biz.GetMeasurement(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
