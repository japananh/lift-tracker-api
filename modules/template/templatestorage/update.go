package templatestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *templatemodel.TemplateUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
