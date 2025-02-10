package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male	Gender = "male"
	Female	Gender = "female"
	Other	Gender = "other"
)

type Role string

const (
	Trainee		Role = "trainee"
	Coach		Role = "coach"
)

type User struct {
	gorm.Model
	Name					string `gorm:"index;not null"`
	Email 					string `gorm:"uniqueIndex;not null"`
	PhoneNumber				string `gorm:"uniqueIndex;not null"`
	Password				string `gorm:"not null"`
	UID  					string `gorm:"not null"`
	EmailVerifiedAt			*time.Time
	PhoneNumberVerifiedAt	*time.Time
	PasswordResetRequested	*time.Time
	Avatar					*string
	Gender 					Gender `gorm:'type:ENUM("male","female","other");default:"male"'`
	Role 					Role `gorm:'type:ENUM("trainee","coach");default:"trainee"'`
	Address 				UserAddress
	DeliveryAddresses		[]DeliveryAddress
	PaymentMethods			[]PaymentMethod
	Banks					[]Bank
	AllowNotifications		bool `gorm:"default:false"`
}
