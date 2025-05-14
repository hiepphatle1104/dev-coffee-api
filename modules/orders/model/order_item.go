package ordermodel

type OrderItem struct {
	OrderID  int `json:"order_id" gorm:"column:order_id;"`
	ItemID   int `json:"item_id" gorm:"column:item_id;"`
	Quantity int `json:"quantity" gorm:"column:quantity;"`
}

func (OrderItem) TableName() string { return "order_items" }
