package model

import "time"

type Order struct {
	ID           uint          `gorm:"primaryKey;autoIncrement;"json:"id"`
	UserID       uint          `gorm:"not null"json:"userID"`   // At the time there is no User entity
	VendorID     uint          `gorm:"not null"json:"vendorID"` // Many-to-One relationship with Vendor
	RegisteredAt time.Time     `gorm:"not null"json:"registeredAt"`
	DeliveryTime time.Duration `gorm:"not null"json:"deliveryTime"`
}
