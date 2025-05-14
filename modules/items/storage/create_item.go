package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *itemmodel.ItemCreation) error {
	err := s.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}
