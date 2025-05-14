package itemservice

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
	"errors"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *itemmodel.ItemCreation) error
}

type CreateItemService struct {
	store CreateItemStorage
}

func NewCreateItemService(store CreateItemStorage) *CreateItemService {
	return &CreateItemService{store: store}
}

func (s *CreateItemService) CreateNewItem(ctx context.Context, data *itemmodel.ItemCreation) error {
	if data.Name == "" {
		return errors.New("name is required")
	}

	if data.UnitPrice < 0 || data.UnitPrice > 500 {
		return errors.New("unit price must be greater than 0")
	}

	return s.store.CreateItem(ctx, data)
}
