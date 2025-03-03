package models

import (
	"time"
)

type PaymentMethod struct {
	ID					string `gorm:"primaryKey"`
	UserID				uint `gorm:"index;not null"`
	LastDigits			string `gorm:"not null"`
	ExpiryMonth			string `gorm:"not null"`
	ExpiryYear			string `gorm:"not null"`
	Brand				string `gorm:"not null"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
}
