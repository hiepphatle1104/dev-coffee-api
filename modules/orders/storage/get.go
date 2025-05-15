package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetOrderByID(ctx context.Context, id int) (*ordermodel.Order, error) {

	var data ordermodel.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&data).Error; err != nil {
			return err
		}

		var orderItems []ordermodel.OrderItem
		err := tx.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", data.ID).Find(&orderItems).Error
		if err != nil {
			return err
		}

		data.OrderItems = &orderItems

		return nil
	})

	return &data, err
}

func (s *sqlStorage) GetOrders(ctx context.Context, paging *ordermodel.Paging) (*[]ordermodel.Order, error) {
	var data []ordermodel.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Offset(paging.Offset()).Limit(paging.Limit).Find(&data).Error; err != nil {
			return err
		}

		for i := range data {
			var orderItems []ordermodel.OrderItem
			err := tx.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", data[i].ID).Find(&orderItems).Error
			if err != nil {
				return err
			}

			data[i].OrderItems = &orderItems
		}

		return nil
	})

	return &data, err
}

func (s *sqlStorage) GetOrderItems(ctx context.Context, orderId int) (*[]ordermodel.OrderItem, error) {
	var orderItems []ordermodel.OrderItem
	err := s.db.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", orderId).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	return &orderItems, nil
}
