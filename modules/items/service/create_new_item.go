package itemservice

import (
	"context"
	"dev-coffee-api/common"
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

	if data.Image != nil {
		err := common.ValidateImage(data.Image)
		if err != nil {
			return err
		}
	}

	return s.store.CreateItem(ctx, data)
}
