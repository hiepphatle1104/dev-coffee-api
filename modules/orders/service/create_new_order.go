package orderservice

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	"errors"
)

type CreateNewOrderService interface {
	Create(ctx context.Context, data *ordermodel.OrderCreation) error
}

type CreateOrderService struct {
	store CreateNewOrderService
}

func NewCreateOrderService(store CreateNewOrderService) *CreateOrderService {
	return &CreateOrderService{store: store}
}

func (s *CreateOrderService) CreateOrder(ctx context.Context, data *ordermodel.OrderCreation) error {
	if data.CustomerName == "" {
		return errors.New("customer name is required")
	}

	return s.store.Create(ctx, data)
}
