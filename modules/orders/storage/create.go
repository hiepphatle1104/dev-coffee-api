package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

func (s *sqlStorage) Create(ctx context.Context, data *ordermodel.OrderCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
