package gincollection

import (
	"errors"
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/collection/collectionbiz"
	"lift-tracker-api/modules/collection/collectionmodel"
	"lift-tracker-api/modules/collection/collectionstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data collectionmodel.CollectionUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if data.ParentId == int(uid.GetLocalID()) {
			panic(common.ErrInvalidRequest(errors.New("invalid request")))
		}
		
		db := appCtx.GetMainDBConnection()
		store := collectionstorage.NewSQLStore(db)
		biz := collectionbiz.NewUpdateCollectionBiz(store)

		if err := biz.UpdateCollection(
			c.Request.Context(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
