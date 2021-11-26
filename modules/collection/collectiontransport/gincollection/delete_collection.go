package gincollection

import (
	"fmt"
	"lift-tracker-api/common"
	"lift-tracker-api/component"
	"lift-tracker-api/modules/collection/collectionbiz"
	"lift-tracker-api/modules/collection/collectionmodel"
	"lift-tracker-api/modules/collection/collectionstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCollection(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type reqData struct {
			FakeIds []string `json:"ids" binding:"required" validate:"required"`
			Ids     []int
		}
		var d reqData

		if err := c.BindJSON(&d); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		d.Ids = make([]int, len(d.FakeIds))

		for i, fakeId := range d.FakeIds {
			uid, err := common.FromBase58(fakeId)
			if err != nil {
				// panic(common.ErrInvalidRequest(err))
				fmt.Println(err)
			}
			d.Ids[i] = int(uid.GetLocalID())
		}

		store := collectionstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := collectionbiz.NewDeleteCollectionBiz(store)

		invalidIds, err := biz.DeleteCollection(c.Request.Context(), d.Ids)
		if err != nil {
			panic(err)
		}

		var fakeInvalidIds []string
		for _, invalidId := range invalidIds {
			uid := common.NewUID(uint32(invalidId), common.DbTypeCollection, 1)
			fakeInvalidIds = append(fakeInvalidIds, uid.String())
		}

		if len(fakeInvalidIds) == 0 {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": fakeInvalidIds,
				"message": fmt.Sprintf(
					"Cannot delete %s/%s ids",
					strconv.Itoa(len(fakeInvalidIds)),
					strconv.Itoa(len(d.FakeIds)),
				),
				"error_key": fmt.Sprintf("Err%sNotFound", collectionmodel.EntityName),
			})
		}
	}
}
