package orderservice

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
)

type GetOrderByIdStorage interface {
	GetOrderByID(ctx context.Context, id int) (*ordermodel.Order, error)
}

type GetOrderByIdService struct {
	store GetOrderByIdStorage
}

func NewGetOrderByIdService(store GetOrderByIdStorage) *GetOrderByIdService {
	return &GetOrderByIdService{store: store}
}

func (s *GetOrderByIdService) GetOrderByID(ctx context.Context, id int) (*ordermodel.Order, error) {
	order, err := s.store.GetOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
