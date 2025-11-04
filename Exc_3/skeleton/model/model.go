package model

import "time"

type Drink struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex;not null"`
	PriceCents int    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	DrinkID   uint `gorm:"not null;index"`
	Qty       int  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Drink Drink `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type TotalledOrder struct {
	OrderID        uint
	DrinkName      string
	Qty            int
	LineTotalCents int
	CreatedAt      time.Time
}
