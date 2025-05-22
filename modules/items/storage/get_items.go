package itemstorage

import (
	"context"
	"dev-coffee-api/common"
	itemmodel "dev-coffee-api/modules/items/model"
)

func (s *sqlStorage) GetItems(ctx context.Context, paging *common.Paging) (*[]itemmodel.Item, error) {
	var items []itemmodel.Item
	if err := s.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &items, nil
}
