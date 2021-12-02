package collectionstorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/collection/collectionmodel"
)

func (s *sqlStore) ListCollectionByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *collectionmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]collectionmodel.Collection, error) {
	db := s.db
	db = db.Table(collectionmodel.Collection{}.TableName()).
		Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CreatedBy > 0 {
			db = db.Where("created_by = ?", v.CreatedBy)
		}
	}

	// db.Count should come before db.Preload
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])

		// if moreKeys[i] == "Collection" {
		// do something (E.g. call api)
		// }
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var result []collectionmodel.Collection

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
