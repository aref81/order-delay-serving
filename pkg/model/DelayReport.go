package model

import "time"

type DelayReport struct {
	ID          uint          `gorm:"primaryKey"`
	OrderID     uint          `gorm:"not null"`
	VendorID    uint          `gorm:"not null"` // Many-to-One relationship with Vendor
	DelayAmount time.Duration `gorm:"not null"`
}
