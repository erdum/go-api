package models

import (
	"time"
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
	ID						uint
	Name					string `gorm:"index;not null" faker:"name"`
	Email 					string `gorm:"uniqueIndex;not null" faker:"email"`
	PhoneNumber				string `gorm:"uniqueIndex;not null" faker:"phone_number"`
	Password				string `gorm:"not null"`
	UID  					string `gorm:"not null" faker:"uuid_digit"`
	EmailVerifiedAt			*time.Time `faker:"-"`
	PhoneNumberVerifiedAt	*time.Time `faker:"-"`
	PasswordResetRequested	*time.Time `faker:"-"`
	Avatar					*string `faker:"-"`
	Gender 					Gender `gorm:'type:ENUM("male","female","other");default:"male"' faker:"-"`
	Role 					Role `gorm:'type:ENUM("trainee","coach");default:"trainee"' faker:"-"`
	Address 				UserAddress `faker:"-"`
	DeliveryAddresses		[]DeliveryAddress `faker:"-"`
	PaymentMethods			[]PaymentMethod `faker:"-"`
	Banks					[]Bank `faker:"-"`
	FcmToken				*string `faker:"-"`
	AllowNotifications		bool `gorm:"default:false" faker:"-"`
	CreatedAt				time.Time `faker:"-"`
	UpdatedAt				time.Time `faker:"-"`
	DeletedAt				*time.Time `faker:"-"`
}
