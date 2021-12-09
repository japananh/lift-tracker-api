package templatestorage

import (
	"context"
	"lift-tracker-api/common"
	"lift-tracker-api/modules/template/templatemodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.
		Table(templatemodel.Template{}.TableName()).
		Delete(&templatemodel.Template{}, "id = ?", id).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
