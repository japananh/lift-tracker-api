package templatestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

func (s *sqlStore) CreateTemplate(ctx context.Context, data *templatemodel.TemplateCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
