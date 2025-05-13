package business

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

type ListDrinkStorage interface {
	List(ctx context.Context, paging *model.Paging) (*[]model.DrinkItem, error)
}

type listDrinkBusiness struct {
	storage ListDrinkStorage
}

func NewListDrinkBusiness(storage ListDrinkStorage) *listDrinkBusiness {
	return &listDrinkBusiness{storage: storage}
}

func (b *listDrinkBusiness) ListDrink(ctx context.Context, paging *model.Paging) (*[]model.DrinkItem, error) {
	data, err := b.storage.List(ctx, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
