package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID				uint `gorm:"uniqueIndex:unique_booking;index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	SessionID			uint `gorm:"uniqueIndex:unique_booking;index;not null"`
	Session				Session `gorm:"foreignKey:SessionID"`
	TransactionID		string `gorm:"uniqueIndex;not null"`
}
