package storage

import (
	"context"
	"dev-coffee-api/modules/drinks/model"
	"fmt"
)

func (s *sqlStorage) Delete(ctx context.Context, id int) error {
	err := s.db.Table(model.DrinkItem{}.TableName()).Where("id = ?", id).Delete(nil).Error
	if err != nil {
		return fmt.Errorf("error while deleting")
	}

	return nil
}
