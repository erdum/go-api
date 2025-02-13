package models

import (
	"time"
)

type CoachCertificate struct {
	Name				string
	Institution			string
	DateOfCompletion	time.Time
	Description			string
}

type CoachDetail struct {
	ID					uint
	CoachID				uint `gorm:"uniqueIndex;not null"`
	Coach				User `gorm:"foreignKey:CoachID"`
	WalletBalance		float32 `gorm:"default:0;not null"`
	Expertise			[]string `gorm:"serializer:json;not null"`
	Experience			uint `gorm:"not null"`
	About				string `gorm:"not null"`
	PerSlotPrice		float32 `gorm:"not null"`
	PerHourPrice		float32 `gorm:"not null"`
	Certificates		[]CoachCertificate `gorm:"serializer:json;not null"`
	MediaLinks			[]Media `gorm:"polymorphic:Model"`
	Reviews				[]Review `gorm:"polymorphic:Model"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
