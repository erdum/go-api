package models

import (
	"time"
)

type Booking struct {
	ID					uint
	UserID				uint `gorm:"uniqueIndex:unique_booking;index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	SessionID			uint `gorm:"uniqueIndex:unique_booking;index;not null"`
	Session				Session `gorm:"foreignKey:SessionID"`
	TransactionID		string `gorm:"uniqueIndex;not null"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
