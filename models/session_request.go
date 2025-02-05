package models

import (
	"gorm.io/gorm"
)

type SessionReqStatus string

const (
	SessionReqPending			SessionReqStatus = "pending"
	SessionReqAccepted			SessionReqStatus = "accepted"
	SessionReqRejected			SessionReqStatus = "rejected"
)

type SessionRequest struct {
	gorm.Model
	SessionID			uint `gorm:"uniqueIndex:single_session_req;index;not null"`
	Session				Session `gorm:"foreignKey:SessionID"`
	UserID				uint `gorm:"uniqueIndex:single_session_req;index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	Status 				SessionReqStatus `gorm:'type:ENUM("pending","accepted","rejected");default:"pending"'`
}
