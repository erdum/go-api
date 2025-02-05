package models

import (
	"time"

	"gorm.io/gorm"
)

type SessionType string

const (
	SessionTypePublic		SessionType = "public"
	SessionTypePrivate		SessionType = "private"
	SessionTypeCustom		SessionType = "custom"
)

type Session struct {
	gorm.Model
	Title				string `gorm:"not null"`
	Description			string `gorm:"not null"`
	Activities			[]string `gorm:"serializer:json;not null"`
	BannerImage			*string
	Amount				float32 `gorm:"not null"`
	Slots				uint `gorm:"not null"`
	Date				time.Time `gorm:"not null"`
	StartTime			time.Time `gorm:"not null"`
	EndTime				time.Time `gorm:"not null"`
	StartedAt			*time.Time
	EndedAt				*time.Time
	Location			string `gorm:"not null"`
	Lat					float32 `gorm:"not null"`
	Long				float32 `gorm:"not null"`
	Type				SessionType `gorm:'type:ENUM("public","private","custom");default:"public"'`
	CoachID				uint `gorm:"index;not null"`
	Coach				User `gorm:"foreignKey:CoachID"`
}
