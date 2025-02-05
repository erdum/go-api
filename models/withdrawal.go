package models

import (
	"gorm.io/gorm"
)

type WithdrawalStatus string

const (
	WithdrawalPending		WithdrawalStatus = "pending"
	WithdrawalApproved		WithdrawalStatus = "approved"
	WithdrawalRejected		WithdrawalStatus = "rejected"
)

type Withdrawal struct {
	gorm.Model
	UserID				uint `gorm:"index;not null"`
	BankID				string `gorm:"index;not null"`
	Amount				float32
	Status				WithdrawalStatus `gorm:'type:ENUM("pending","approved","rejected");default:"pending"'`
	RejectionReason		*string
}
