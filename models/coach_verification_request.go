package models

import (
	"gorm.io/gorm"
)

type CoachVerificationStatus string

const (
	CoachVerificationPending		CoachVerificationStatus = "pending"
	CoachVerificationAccepted		CoachVerificationStatus = "accepted"
	CoachVerificationRejected		CoachVerificationStatus = "rejected"
)

type CoachVerificationRequest struct {
	gorm.Model
	CoachID					uint `gorm:"uniqueIndex;not null"`
	Coach					User `gorm:"foreignKey:CoachID"`
	Status 					CoachVerificationStatus `gorm:'type:ENUM("pending","accepted","rejected");default:"pending"'`
	DrivingLicenseFront		Media `gorm:"polymorphic:Model;not null"`
	DrivingLicenseBack		Media `gorm:"polymorphic:Model;not null"`
	Passport				Media `gorm:"polymorphic:Model;not null"`
}
