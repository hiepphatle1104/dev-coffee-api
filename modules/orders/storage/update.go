package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {

	return s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(ordermodel.OrderUpdate{}.TableName()).Where("id = ?", id).Updates(&data).Error
		if err != nil {
			return err
		}

		if err = tx.Where("order_id = ?", id).Delete(&ordermodel.OrderItem{}).Error; err != nil {
			return err
		}

		for _, orderItem := range *data.OrderItems {
			orderItem.OrderID = id
			if err = tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
