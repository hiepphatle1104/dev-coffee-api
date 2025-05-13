package business

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

type GetDrinkStorage interface {
	GetByID(ctx context.Context, id int) (*model.DrinkItem, error)
}

type getDrinkBusiness struct {
	storage GetDrinkStorage
}

func NewGetDrinkBusiness(storage GetDrinkStorage) *getDrinkBusiness {
	return &getDrinkBusiness{storage: storage}
}

func (b *getDrinkBusiness) GetItemById(ctx context.Context, id int) (*model.DrinkItem, error) {
	data, err := b.storage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
