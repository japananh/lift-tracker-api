package gincollection

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/collection/collectionbiz"
	"lift-tracker-api/modules/collection/collectionstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := collectionstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := collectionbiz.NewCollectionBiz(store)

		data, err := biz.GetCollection(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
