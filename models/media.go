package models

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	URL					string
	ModelID				uint
	ModelType			string
}
