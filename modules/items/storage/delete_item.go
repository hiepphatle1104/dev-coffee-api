package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) DeleteItemById(ctx context.Context, id int) error {
	if err := s.db.Table(itemmodel.Item{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
