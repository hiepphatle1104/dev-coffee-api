package orderstorage

import (
	"context"
	"dev-coffee-api/common"
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

func (s *sqlStorage) GetOrders(ctx context.Context, paging *common.Paging) (*[]ordermodel.Order, error) {
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
func (s *sqlStorage) GetOrderItemsByID(ctx context.Context, id int) (*[]ordermodel.OrderItem, error) {
	var orderItems []ordermodel.OrderItem
	if err := s.db.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", id).Find(&orderItems).Error; err != nil {
		return nil, err
	}

	return &orderItems, nil
}

func (s *sqlStorage) GetOrderItems(ctx context.Context, id int) (*[]ordermodel.MergedOrderItem, error) {
	var mergedOrderItems []ordermodel.MergedOrderItem

	err := s.db.Table("order_items").
		Joins("JOIN items ON order_items.item_id = items.id").
		Where("order_id = ?", id).
		Scan(&mergedOrderItems).Error

	if err != nil {
		return nil, err
	}

	return &mergedOrderItems, nil
}
