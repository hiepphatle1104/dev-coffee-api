package storage

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

func (s *sqlStorage) Update(ctx context.Context, id int, data *model.DrinkItemUpdate) error {
	err := s.db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}
