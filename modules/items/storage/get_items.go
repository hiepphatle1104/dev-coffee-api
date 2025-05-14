package itemstorage

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) GetItems(ctx context.Context, paging *itemmodel.Paging) (*[]itemmodel.Item, error) {
	var items []itemmodel.Item
	if err := s.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &items, nil
}
