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

func ListCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter collectionmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if filter.FakeCreatedBy != "" {
			createdByUID, err := common.FromBase58(filter.FakeCreatedBy)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.CreatedBy = int(createdByUID.GetLocalID())
		}

		if filter.FakeParentId != "" {
			parentUID, err := common.FromBase58(filter.FakeParentId)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.ParentId = int(parentUID.GetLocalID())
		}

		var paging common.Paging

		paging.Fulfill()

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := collectionstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := collectionbiz.NewListCollectionBiz(store)

		result, err := biz.ListCollection(c.Request.Context(), &filter, &paging)
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
