package paymentmodel

import ordermodel "dev-coffee-api/modules/orders/model"

type Payment struct {
	ID      int               `json:"-" gorm:"column:id;"`
	OrderID int               `json:"order_id" gorm:"column:order_id;"`
	Amount  float64           `json:"amount" gorm:"column:amount;"`
	Method  *PaymentMethod    `json:"method" gorm:"column:method;"`
	PaidAt  string            `json:"paid_at" gorm:"column:paid_at;"`
	Order   *ordermodel.Order `json:"-" gorm:"foreignKey:OrderID;references:ID"`
}

func (Payment) TableName() string { return "payments" }

type PaymentCreation struct {
	OrderID int            `json:"order_id" gorm:"column:order_id;"`
	Amount  float64        `json:"-" gorm:"column:amount;"`
	Method  *PaymentMethod `json:"method" gorm:"column:method;"`
}

func (PaymentCreation) TableName() string { return Payment{}.TableName() }
