package itemservice

import (
	"context"
	"dev-coffee-api/common"
	itemmodel "dev-coffee-api/modules/items/model"
)

type UpdateItemByIdStorage interface {
	UpdateItemById(ctx context.Context, id int, data *itemmodel.ItemUpdate) error
}

type UpdateItemByIdService struct {
	store UpdateItemByIdStorage
}

func NewUpdateItemByIdService(store UpdateItemByIdStorage) *UpdateItemByIdService {
	return &UpdateItemByIdService{store: store}
}

func (s *UpdateItemByIdService) UpdateItemById(ctx context.Context, id int, data *itemmodel.ItemUpdate) error {
	if data.Image != nil {
		err := common.ValidateImage(data.Image)
		if err != nil {
			return err
		}
	}

	return s.store.UpdateItemById(ctx, id, data)
}
