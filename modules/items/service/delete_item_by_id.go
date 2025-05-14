package itemservice

import "context"

type DeleteItemByIdStorage interface {
	DeleteItemById(ctx context.Context, id int) error
}

type DeleteItemByIdService struct {
	store DeleteItemByIdStorage
}

func NewDeleteItemByIdService(store DeleteItemByIdStorage) *DeleteItemByIdService {
	return &DeleteItemByIdService{store: store}
}

func (s *DeleteItemByIdService) DeleteItemById(ctx context.Context, id int) error {
	return s.store.DeleteItemById(ctx, id)
}
