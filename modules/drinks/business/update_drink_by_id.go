package business

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

type UpdateDrinkStorage interface {
	Update(ctx context.Context, id int, data *model.DrinkItemUpdate) error
}

type updateDrinkBusiness struct {
	storage UpdateDrinkStorage
}

func NewUpdateDrinkBusiness(storage UpdateDrinkStorage) *updateDrinkBusiness {
	return &updateDrinkBusiness{storage: storage}
}

func (b *updateDrinkBusiness) UpdateDrink(ctx context.Context, id int, data *model.DrinkItemUpdate) error {
	err := b.storage.Update(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}
