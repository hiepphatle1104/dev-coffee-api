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

func (s *UpdateOrderByIdService) UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {

	if data.CustomerName == "" {
		return errors.New("customer name is required")
	}

	if data.OrderItems == nil || len(*data.OrderItems) == 0 {
		return errors.New("order items is required")
	}

	for i := range *data.OrderItems {
		orderItem := &(*data.OrderItems)[i]
		item, err := s.itemStore.GetItemById(ctx, orderItem.ItemID)
		if err != nil {
			return errors.New("item not found")
		}

		if !*item.Available {
			return errors.New("item is not available")
		}

		orderItem.UnitPrice = item.UnitPrice
	}

	return s.store.UpdateOrderByID(ctx, id, data)
}
