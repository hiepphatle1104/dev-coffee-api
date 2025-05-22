package orderservice

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

type GetOrderItemStorage interface {
	GetOrderItems(ctx context.Context, id int) (*[]ordermodel.MergedOrderItem, error)
}

type GetOrderItemsService struct {
	store GetOrderItemStorage
}

func NewGetOrderItemsService(store GetOrderItemStorage) *GetOrderItemsService {
	return &GetOrderItemsService{store: store}
}

func (s *GetOrderItemsService) GetOrderItems(ctx context.Context, id int) (*[]ordermodel.MergedOrderItem, error) {
	return s.store.GetOrderItems(ctx, id)
}
