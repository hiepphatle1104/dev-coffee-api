package business

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
	"fmt"
)

type CreateDrinkStorage interface {
	Create(ctx context.Context, data *model.DrinkItemCreation) error
}

type createDrinkBusiness struct {
	storage CreateDrinkStorage
}

func NewCreateDrinkBusiness(storage CreateDrinkStorage) *createDrinkBusiness {
	return &createDrinkBusiness{storage: storage}
}

func (b *createDrinkBusiness) CreateNewDrink(ctx context.Context, data *model.DrinkItemCreation) error {
	var name string = data.Name
	var price float64 = data.Price

	if name == "" {
		return fmt.Errorf("name is required")
	}

	if price < 0 {
		return fmt.Errorf("price must be greater than or equal to 0")
	}

	err := b.storage.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
