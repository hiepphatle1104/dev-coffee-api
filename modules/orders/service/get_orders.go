package orderservice

import (
	"context"
	"dev-coffee-api/common"
	ordermodel "dev-coffee-api/modules/orders/model"
)

type GetOrdersStorage interface {
	GetOrders(ctx context.Context, paging *common.Paging) (*[]ordermodel.Order, error)
}

type GetOrdersService struct {
	store GetOrdersStorage
}

func NewGetOrdersService(store GetOrdersStorage) *GetOrdersService {
	return &GetOrdersService{store: store}
}

func (s *GetOrdersService) GetOrders(ctx context.Context, paging *common.Paging) (*[]ordermodel.Order, error) {
	orders, err := s.store.GetOrders(ctx, paging)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
