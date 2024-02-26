package model

import "time"

type DelayReport struct {
	ID          uint          `gorm:"primaryKey;autoIncrement;"json:"id"`
	OrderID     uint          `gorm:"not null"json:"orderID"`
	VendorID    uint          `gorm:"not null"json:"vendorID"` // Many-to-One relationship with Vendor
	DelayAmount time.Duration `gorm:"not null"json:"delayAmount"`
}
