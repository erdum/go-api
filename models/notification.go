package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID				uint `gorm:"index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Title				string `gorm:"not null"`
	Body				string `gorm:"not null"`
	Image				*string
	Data				string `gorm:"serializer:json"`
}
