package orderservice

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

type UpdateOrderByIdStorage interface {
	UpdateOrderByID(ctx context.Context, id int, data *ordermodel.OrderUpdate) error
}

type UpdateOrderByIdService struct {
	store UpdateOrderByIdStorage
}

func NewUpdateOrderByIdService(store UpdateOrderByIdStorage) *UpdateOrderByIdService {
	return &UpdateOrderByIdService{store: store}
}

func (s *UpdateOrderByIdService) UpdateOrderById(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	return s.store.UpdateOrderByID(ctx, id, data)
}
