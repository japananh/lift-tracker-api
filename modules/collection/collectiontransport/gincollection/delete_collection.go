package gincollection

import (
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/collection/collectionbiz"
	"lift-tracker-api/modules/collection/collectionstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := collectionstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := collectionbiz.NewDeleteCollectionBiz(store)

		if err := biz.DeleteCollection(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
