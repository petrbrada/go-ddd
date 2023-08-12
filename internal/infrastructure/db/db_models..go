package db

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Price    float64
	SellerID uuid.UUID `gorm:"index"`
}

type Seller struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}