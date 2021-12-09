package templatestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

func (s *sqlStore) ListTemplateByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *templatemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]templatemodel.Template, error) {
	db := s.db
	db = db.Table(templatemodel.Template{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.CreatedBy >= 1 {
			db = db.Where("created_by = ?", v.CreatedBy)
		}

		if v.CollectionId >= 1 {
			db = db.Where("mechanics = ?", v.CollectionId)
		}
	}

	// db.Count should come before db.Preload
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var result []templatemodel.Template

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
