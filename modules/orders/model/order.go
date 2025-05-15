package ordermodel

import "context"

type Order struct {
	ID           int          `json:"id" gorm:"column:id;"`
	CustomerName string       `json:"customer_name" gorm:"column:customer_name;"`
	Status       *OrderStatus `json:"status" gorm:"column:status;"`
	OrderItems   *[]OrderItem `json:"order_items" gorm:"-"`
	CreatedAt    string       `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    string       `json:"updated_at" gorm:"column:updated_at;"`
}

func (Order) TableName() string { return "orders" }

type OrderCreation struct {
	ID           int          `json:"-" gorm:"column:id;"`
	CustomerName string       `json:"customer_name" gorm:"column:customer_name;"`
	OrderItems   *[]OrderItem `json:"order_items" gorm:"-"`
}

func (OrderCreation) TableName() string { return Order{}.TableName() }

type OrderUpdate struct {
	ID           int          `json:"-" gorm:"column:id;"`
	CustomerName string       `json:"customer_name" gorm:"column:customer_name;"`
	Status       *OrderStatus `json:"status" gorm:"column:status;"`
	OrderItems   *[]OrderItem `json:"order_items" gorm:"-"`
}

func (OrderUpdate) TableName() string { return Order{}.TableName() }

func NewOrderUpdateStatus(status *OrderStatus) *OrderUpdate {
	return &OrderUpdate{Status: status}
}

type OrderStorage interface {
	GetOrderByID(ctx context.Context, id int) (*Order, error)
}
