package models

import "gorm.io/gorm"

type Gender string

const (
	Male	Gender = "male"
	Female	Gender = "female"
	Other	Gender = "other"
)

type User struct {
	gorm.Model
	Name				string `gorm:"index"`
	Email 				string `gorm:"index:idx_email"`
	PhoneNumber			string `gorm:"index:idx_phone"`
	Avatar				*string
	Gender Gender `gorm:'type:ENUM("male","female","other");default:"other"'`
	UserPreferenceID	uint `gorm:"index:idx_user_preference_id"`
	Preference 			UserPreference `gorm:"foreignKey:UserPreferenceID"`
}
