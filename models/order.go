package models

import (
	"time"
)

type OrderStatus string

const (
	OrderInProgress			OrderStatus = "inprogress"
	OrderCompleted			OrderStatus = "completed"
	OrderCanceled			OrderStatus = "canceled"
)

type TrackingStatus string

const (
	TrackingStatusConfirmed				TrackingStatus = "confirmed"
	TrackingStatusInProgress			TrackingStatus = "inprogress"
	TrackingStatusShipped				TrackingStatus = "shipped"
	TrackingStatusOutForDelivery		TrackingStatus = "out for delivery"
	TrackingStatusDelivered				TrackingStatus = "delivered"
)

type Order struct {
	ID					uint
	UserID				uint `gorm:"index;not null"`
	User				User `gorm:"foreignKey:UserID"`
	DeliveryAddressID	uint `gorm:"index;not null"`
	DeliveryAddress		DeliveryAddress `gorm:"foreignKey:DeliveryAddressID;not null"`
	PaymentMethodID		string `gorm:"index;not null"`
	PaymentMethod		PaymentMethod `gorm:"foreignKey:PaymentMethodID;not null"`
	Products			[]Product `gorm:"many2many:order_products;not null"`
	Status 				OrderStatus `gorm:'type:ENUM("inprogress","completed","canceled");default:"inprogress"'`
	TrackingStatus		TrackingStatus `gorm:'type:ENUM("confirmed","inprogress","shipped","out for delivery","delivered");default:"inprogress"'`
	TransactionID		string `gorm:"uniqueIndex;not null"`
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
