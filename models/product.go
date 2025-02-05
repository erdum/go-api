package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CoachID				uint `gorm:"index;not null"`
	Coach				User `gorm:"foreignKey:CoachID"`
	Title				string `gorm:"not null"`
	Description			string `gorm:"not null"`
	Amount				float32 `gorm:"not null"`
	Quantity			uint `gorm:"not null"`
	Colors				[]string `gorm:"serializer:json"`
	Sizes				[]string `gorm:"serializer:json"`
	Images				[]Media `gorm:"polymorphic:Model"`
	Reviews				[]Review `gorm:"polymorphic:Model"`
}
