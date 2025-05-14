package itemmodel

import "context"

type ItemStore interface {
	GetItemById(ctx context.Context, id int) (*Item, error)
}
