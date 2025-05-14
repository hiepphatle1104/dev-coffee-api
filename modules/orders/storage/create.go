package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) Create(ctx context.Context, data *ordermodel.OrderCreation) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Insert order
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		// Insert order items
		for _, orderItem := range *data.OrderItems {
			orderItem.OrderID = data.ID
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
