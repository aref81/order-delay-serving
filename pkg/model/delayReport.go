package model

import "time"

type DelayReport struct {
	ID          uint          `gorm:"primaryKey;autoIncrement;"json:"id"`
	OrderID     uint          `gorm:"not null"json:"orderID"`
	VendorID    uint          `gorm:"not null"json:"vendorID"` // Many-to-One relationship with Vendor
	DelayAmount time.Duration `gorm:"not null"json:"delayAmount"`
	IssuedAt    time.Time     `gorm:"not null;index'"json:"issuedAt"` // indexed in order to enhance the summary query performance
}

type VendorDelaySummary struct {
	VendorID         uint  `gorm:"column:vendor_id"`
	TotalDelayAmount int64 `gorm:"column:total_delay_amount"`
}
