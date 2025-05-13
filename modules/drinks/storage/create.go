package storage

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

func (s *sqlStorage) Create(ctx context.Context, data *model.DrinkItemCreation) error {
	err := s.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
