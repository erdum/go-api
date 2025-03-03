package models

import (
	"time"
)

type Media struct {
	ID					uint
	URL					string
	ModelID				uint
	ModelType			string
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
