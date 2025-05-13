package storage

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

func (s *sqlStorage) GetByID(ctx context.Context, id int) (*model.DrinkItem, error) {
	var data model.DrinkItem
	err := s.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
