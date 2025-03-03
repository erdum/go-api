package models

import (
	"time"
)

type UserAddress struct {
	ID					uint
	UserID				uint `gorm:"uniqueIndex;not null"`
	Country				string `gorm:"not null"`
	State				string `gorm:"not null"`
	City				string `gorm:"not null"`
	ZipCode				string `gorm:"not null"`
	Address				string `gorm:"not null"`
	Lat					float32 `gorm:"not null"`
	Long				float32 `gorm:"not null"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
