package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID				uint `gorm:"uniqueIndex:single_review;index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Rating				uint8 `gorm:"not null"`
	Text				string `gorm:"not null"`
	ModelID				uint `gorm:"uniqueIndex:single_review;not null"`
	ModelType			string `gorm:"uniqueIndex:single_review;not null"`
}
