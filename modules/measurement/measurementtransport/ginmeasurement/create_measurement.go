package ginmeasurement

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/measurement/measurementbiz"
	"lift-tracker-api/modules/measurement/measurementmodel"
	"lift-tracker-api/modules/measurement/measurementstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMeasurement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data measurementmodel.MeasurementCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := measurementstorage.NewSQLStore(db)
		biz := measurementbiz.NewCreateMeasurementBiz(store)

		if err := biz.CreateMeasurement(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
