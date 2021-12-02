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

func ListMeasurement(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter measurementmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if filter.FakeUserId != "" {
			UserUID, err := common.FromBase58(filter.FakeUserId)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.UserId = int(UserUID.GetLocalID())
		}

		var paging common.Paging

		paging.Fulfill()

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := measurementstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := measurementbiz.NewListMeasurementBiz(store)

		result, err := biz.ListMeasurement(c.Request.Context(), &filter, &paging)
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
