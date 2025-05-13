package model

import "time"

type DrinkItem struct {
	ID        int        `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	Type      string     `json:"type" gorm:"column:type;"`
	Price     float64    `json:"price" gorm:"column:price;"`
	Available int        `json:"available" gorm:"column:available;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;omitempty"`
}

func (DrinkItem) TableName() string { return "drinks" }

type DrinkItemCreation struct {
	ID    int     `json:"-" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Type  string  `json:"type" gorm:"column:type;"`
	Price float64 `json:"price" gorm:"column:price;"`
}

func (DrinkItemCreation) TableName() string { return DrinkItem{}.TableName() }

type DrinkItemUpdate struct {
	Name      string   `json:"name" gorm:"column:name;"`
	Type      string   `json:"type" gorm:"column:type;"`
	Price     *float64 `json:"price" gorm:"column:price;"`
	Available *int     `json:"available" gorm:"column:available;"`
}

func (DrinkItemUpdate) TableName() string { return DrinkItem{}.TableName() }
