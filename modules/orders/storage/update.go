package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		var existOrderItem []ordermodel.OrderItem
		err := tx.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", id).Find(&existOrderItem).Error
		if err != nil {
			return err
		}

		// TODO: Need to fix here

		err = s.db.Table(ordermodel.OrderUpdate{}.TableName()).Where("id = ?", id).Updates(&data).Error
		if err != nil {
			return err
		}
		return nil
	})
}
