package orderstorage

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

func (s *sqlStorage) GetOrderByID(ctx context.Context, id int) (*ordermodel.Order, error) {
	var data ordermodel.Order
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStorage) GetOrders(ctx context.Context, paging *ordermodel.Paging) (*[]ordermodel.Order, error) {
	var data []ordermodel.Order
	if err := s.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
