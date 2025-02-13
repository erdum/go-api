package models

import (
	"time"
)

type Notification struct {
	ID					uint
	UserID				uint `gorm:"index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Title				string `gorm:"not null"`
	Body				string `gorm:"not null"`
	Image				*string
	Data				string `gorm:"serializer:json"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
