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
