package itemservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

type GetItemByIdStorage interface {
	GetItemById(ctx context.Context, id int) (*itemmodel.Item, error)
}

type GetItemByIdService struct {
	store GetItemByIdStorage
}

func NewGetItemByIdService(store GetItemByIdStorage) *GetItemByIdService {
	return &GetItemByIdService{store: store}
}

func (s *GetItemByIdService) GetItemById(ctx context.Context, id int) (*itemmodel.Item, error) {
	item, err := s.store.GetItemById(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}
