package itemservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
	"errors"
)

type DeleteItemByIdStorage interface {
	DeleteItemById(ctx context.Context, id int) error
	GetItemById(ctx context.Context, id int) (*itemmodel.Item, error)
}

type DeleteItemByIdService struct {
	store DeleteItemByIdStorage
}

func NewDeleteItemByIdService(store DeleteItemByIdStorage) *DeleteItemByIdService {
	return &DeleteItemByIdService{store: store}
}

func (s *DeleteItemByIdService) DeleteItemById(ctx context.Context, id int) error {
	exist, err := s.store.GetItemById(ctx, id)
	if err != nil {
		return err
	}

	if exist == nil {
		return errors.New("item not found")
	}
	return s.store.DeleteItemById(ctx, id)
}
