package models

import (
	"gorm.io/gorm"
)

type FitnessGoal string

type userPreferenceTypes struct {
	WeightLoss				string
	StrengthBuilding		string
	FlexibilityImprovement	string
	Beginner				string
	Intermediate			string
	Advance					string
	Yoga					string
	Cardio					string
	StrengthTraining		string
}

var enum = userPreferenceTypes{
	WeightLoss:				"weight_loss",
	StrengthBuilding:		"strength_building",
	FlexibilityImprovement:	"flexibility_improvement",
	Yoga:					"yoga",
	Cardio:					"cardio",
	StrengthTraining:		"strength_training",
}

func getUserPreferenceTypes() userPreferenceTypes {
	return enum
}

type UserPreference struct {
	gorm.Model
	City			string
	Country			string
	State			string
	ZipCode			string
	Address			string
	FitnessGoal 	[]string `gorm:"serializer:json"`
	FitnessLevel 	[]string `gorm:"serializer:json"`
	ActivityChoice	[]string `gorm:"serializer:json"`
}
