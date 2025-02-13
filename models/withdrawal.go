package models

import (
	"time"
)

type WithdrawalStatus string

const (
	WithdrawalPending		WithdrawalStatus = "pending"
	WithdrawalApproved		WithdrawalStatus = "approved"
	WithdrawalRejected		WithdrawalStatus = "rejected"
)

type Withdrawal struct {
	ID					uint
	UserID				uint `gorm:"index;not null"`
	BankID				string `gorm:"index;not null"`
	Amount				float32
	Status				WithdrawalStatus `gorm:'type:ENUM("pending","approved","rejected");default:"pending"'`
	RejectionReason		*string
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
