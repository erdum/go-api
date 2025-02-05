package models

import (
	"gorm.io/gorm"
)

type traineePreferences struct {
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

var enum = traineePreferences{
	WeightLoss:				"weight_loss",
	StrengthBuilding:		"strength_building",
	FlexibilityImprovement:	"flexibility_improvement",
	Yoga:					"yoga",
	Cardio:					"cardio",
	StrengthTraining:		"strength_training",
}

func getTraineePreferenceTypes() traineePreferences {
	return enum
}

type TraineeDetail struct {
	gorm.Model
	UserID				uint `gorm:"uniqueIndex;not null"`
	User				User `gorm:"foreignKey:UserID"`
	FitnessGoal 		[]string `gorm:"serializer:json;not null"`
	FitnessLevel 		[]string `gorm:"serializer:json;not null"`
	ActivityChoice		[]string `gorm:"serializer:json;not null"`
}
