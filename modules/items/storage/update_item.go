package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) UpdateItemById(ctx context.Context, id int, data *itemmodel.ItemUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
