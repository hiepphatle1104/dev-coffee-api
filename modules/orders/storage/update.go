package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

func (s *sqlStorage) UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	if err := s.db.Model(&data).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
