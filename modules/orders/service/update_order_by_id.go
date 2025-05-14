package orderservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
	ordermodel "dev-coffee-api/modules/orders/model"
	"errors"
)

type UpdateOrderByIdStorage interface {
	UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error
}

type UpdateOrderByIdService struct {
	store     UpdateOrderByIdStorage
	itemStore itemmodel.ItemStore
}

func NewUpdateOrderByIdService(store UpdateOrderByIdStorage, itemStore itemmodel.ItemStore) *UpdateOrderByIdService {
	return &UpdateOrderByIdService{store: store, itemStore: itemStore}
}

func (s *UpdateOrderByIdService) UpdateOrderById(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {

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

	return s.store.UpdateOrderByID(ctx, id, data)
}
