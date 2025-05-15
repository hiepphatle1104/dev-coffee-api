package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

func (s *sqlStorage) DeleteOrderById(ctx context.Context, id int) error {
	err := s.db.Table(ordermodel.Order{}.TableName()).Where("id = ?", id).Delete(nil).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStorage) DeleteOrderItemById(ctx context.Context, orderId, id int) error {
	err := s.db.Table(ordermodel.OrderItem{}.TableName()).Where("orderId = ?", orderId).Where("id = ?", id).Delete(nil).Error
	if err != nil {
		return err
	}

	return nil
}
