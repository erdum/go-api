package models

import (
	"time"
)

type Bank struct {
	ID					string `gorm:"primaryKey"`
	UserID				uint `gorm:"index;not null"`
	HolderName			string `gorm:"not null"`
	LastDigits			string `gorm:"not null"`
	BankName			string `gorm:"not null"`
	RoutingNumberf		string `gorm:"not null"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
}
