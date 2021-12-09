package templatestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindTemplateByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*templatemodel.Template, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var template templatemodel.Template

	if err := db.
		Where(conditions).
		First(&template).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &template, nil
}
