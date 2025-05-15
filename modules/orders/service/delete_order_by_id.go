package orderservice

import "context"

type DeleteOrderByIdStorage interface {
	DeleteOrderById(ctx context.Context, id int) error
}

type DeleteOrderByIdService struct {
	store DeleteOrderByIdStorage
}

func NewDeleteOrderByIdService(store DeleteOrderByIdStorage) *DeleteOrderByIdService {
	return &DeleteOrderByIdService{store: store}
}

func (s *DeleteOrderByIdService) DeleteOrderByID(ctx context.Context, id int) error {
	return s.store.DeleteOrderById(ctx, id)
}
