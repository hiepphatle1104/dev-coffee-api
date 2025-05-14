package orderservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
	ordermodel "dev-coffee-api/modules/orders/model"
	"errors"
)

type CreateNewOrderStorage interface {
	Create(ctx context.Context, data *ordermodel.OrderCreation) error
}

type CreateOrderService struct {
	store     CreateNewOrderStorage
	itemStore itemmodel.ItemStore
}

func NewCreateOrderService(store CreateNewOrderStorage, itemStore itemmodel.ItemStore) *CreateOrderService {
	return &CreateOrderService{store: store, itemStore: itemStore}
}

func (s *CreateOrderService) CreateOrder(ctx context.Context, data *ordermodel.OrderCreation) error {
	if data.CustomerName == "" {
		return errors.New("customer name is required")
	}

	if data.OrderItems == nil || len(*data.OrderItems) == 0 {
		return errors.New("order items is required")
	}

	for _, orderItem := range *data.OrderItems {
		_, err := s.itemStore.GetItemById(ctx, orderItem.ItemID)
		if err != nil {
			return errors.New("item not found")
		}
	}

	return s.store.Create(ctx, data)
}
