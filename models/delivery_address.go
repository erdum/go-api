package models

import (
	"gorm.io/gorm"
)

type DeliveryAddress struct {
	gorm.Model
	UserID				uint `gorm:"index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Country				string `gorm:"not null"`
	State				string `gorm:"not null"`
	City				string `gorm:"not null"`
	ZipCode				string `gorm:"not null"`
	Address				string `gorm:"not null"`
	Lat					float32 `gorm:"not null"`
	Long				float32 `gorm:"not null"`
}
