package storage

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
)

func (s *sqlStorage) List(ctx context.Context, paging *model.Paging) (*[]model.DrinkItem, error) {
	var drinks []model.DrinkItem
	err := s.db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&drinks).Error
	if err != nil {
		return nil, err
	}

	return &drinks, nil
}
