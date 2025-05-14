package itemservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

type GetItemListStorage interface {
	GetItems(ctx context.Context, paging *itemmodel.Paging) (*[]itemmodel.Item, error)
}

type GetItemListService struct {
	store GetItemListStorage
}

func NewGetItemListService(store GetItemListStorage) *GetItemListService {
	return &GetItemListService{store: store}
}

func (s *GetItemListService) GetItemList(ctx context.Context, paging *itemmodel.Paging) (*[]itemmodel.Item, error) {
	items, err := s.store.GetItems(ctx, paging)
	if err != nil {
		return nil, err
	}

	return items, nil
}
