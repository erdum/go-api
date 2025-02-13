package models

import (
	"time"
)

type Cart struct {
	ID					uint
	UserID				uint `gorm:"uniqueIndex;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Products			[]Product `gorm:"many2many:cart_products"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
