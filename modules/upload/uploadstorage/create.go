package uploadstorage

import (
	"context"
	"lift-tracker-api/common"
)

func (store *sqlStore) Create(context context.Context, data *common.Image) error {
	db := store.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
