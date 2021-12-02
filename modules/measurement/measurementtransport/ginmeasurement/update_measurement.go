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

func UpdateMeasurement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data measurementmodel.MeasurementUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// if data.ParentId == int(uid.GetLocalID()) {
		// 	panic(common.ErrInvalidRequest(errors.New("invalid request")))
		// }

		db := appCtx.GetMainDBConnection()
		store := measurementstorage.NewSQLStore(db)
		biz := measurementbiz.NewUpdateMeasurementBiz(store)

		if err := biz.UpdateMeasurement(
			c.Request.Context(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
