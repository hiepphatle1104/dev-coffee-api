package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) GetItemById(ctx context.Context, id int) (*itemmodel.Item, error) {
	var item itemmodel.Item
	if err := s.db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
