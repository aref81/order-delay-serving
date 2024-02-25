package model

import "time"

type Order struct {
	ID           uint          `gorm:"primaryKey"`
	UserID       uint          `gorm:"not null"` // At the time there is no User entity
	VendorID     uint          `gorm:"not null"` // Many-to-One relationship with Vendor
	RegisteredAt time.Time     `gorm:"not null"`
	DeliveryTime time.Duration `gorm:"not null"`
}
