package gincollection

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/collection/collectionbiz"
	"lift-tracker-api/modules/collection/collectionmodel"
	"lift-tracker-api/modules/collection/collectionstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data collectionmodel.CollectionCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId, convertIdErr := common.FromBase58(data.UserId)
		if convertIdErr != nil {
			panic(common.ErrInvalidRequest(convertIdErr))
		}

		data.CreatedBy = int(userId.GetLocalID())

		db := appCtx.GetMainDBConnection()
		store := collectionstorage.NewSQLStore(db)
		biz := collectionbiz.NewCreateCollectionBiz(store)

		if err := biz.CreateCollection(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
