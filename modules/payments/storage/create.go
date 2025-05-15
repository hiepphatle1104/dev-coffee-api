package paymentstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	paymentmodel "dev-coffee-api/modules/payments/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) CreateNewPayment(ctx context.Context, data *paymentmodel.PaymentCreation) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Create a new payment
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		// Update order status
		if err := tx.Table(ordermodel.Order{}.
			TableName()).
			Where("id = ?", data.OrderID).
			Update("status", ordermodel.OrderStatusCompleted).Error; err != nil {
			return err
		}

		return nil
	})
}
