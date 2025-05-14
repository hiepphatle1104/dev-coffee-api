package itemmodel

import "time"

type Item struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Type      *ItemType  `json:"type" gorm:"column:type;"`
	UnitPrice float64    `json:"unit_price" gorm:"column:unit_price;"`
	Available *bool      `json:"available" gorm:"column:available;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Item) TableName() string {
	return "items"
}

type ItemCreation struct {
	Name      string    `json:"name" gorm:"column:name;"`
	Type      *ItemType `json:"type" gorm:"column:type;"`
	UnitPrice float64   `json:"unit_price" gorm:"column:unit_price;"`
	Available *bool     `json:"available" gorm:"column:available;"`
}

func (ItemCreation) TableName() string {
	return Item{}.TableName()
}

type ItemUpdate struct {
	Name      string    `json:"name" gorm:"column:name;"`
	Type      *ItemType `json:"type" gorm:"column:type;"`
	UnitPrice float64   `json:"unit_price" gorm:"column:unit_price;"`
	Available *bool     `json:"available" gorm:"column:available;"`
}

func (ItemUpdate) TableName() string {
	return Item{}.TableName()
}
