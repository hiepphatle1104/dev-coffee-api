package business

import (
	"context"
)

type DeleteDrinkStorage interface {
	Delete(ctx context.Context, id int) error
}

type deleteDrinkBusiness struct {
	storage DeleteDrinkStorage
}

func NewDeleteDrinkBusiness(storage DeleteDrinkStorage) *deleteDrinkBusiness {
	return &deleteDrinkBusiness{storage: storage}
}

func (b *deleteDrinkBusiness) DeleteDrink(ctx context.Context, id int) error {
	err := b.storage.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
