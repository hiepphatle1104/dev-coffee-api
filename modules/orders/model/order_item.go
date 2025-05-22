package ordermodel

import (
	"context"
	itemmodel "dev-coffee-api/modules/items/model"
)

type OrderItem struct {
	OrderID   int             `json:"-" gorm:"column:order_id;"`
	ItemID    int             `json:"item_id" gorm:"column:item_id;"`
	Quantity  int             `json:"quantity" gorm:"column:quantity;"`
	UnitPrice float64         `json:"unit_price" gorm:"column:unit_price;"`
	Order     *Order          `json:"-" gorm:"foreignKey:OrderID;references:ID"`
	Item      *itemmodel.Item `json:"-" gorm:"foreignKey:ItemID;references:ID"`
}

func (OrderItem) TableName() string { return "order_items" }

type OrderItemStorage interface {
	GetOrderItemsByID(ctx context.Context, id int) (*[]OrderItem, error)
}

type MergedOrderItem struct {
	ItemID    int     `json:"item_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Name      string  `json:"name"`
}
