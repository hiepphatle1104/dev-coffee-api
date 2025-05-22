package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateItemById(ctx context.Context, id int, data *itemmodel.ItemUpdate) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		exists, err := s.GetItemById(ctx, id)
		if err != nil {
			return err
		}
		if exists == nil {
			return errors.New("item not found")
		}

		if err = s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
			return err
		}
		return nil
	})
}
