package model

type Agent struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"` // At the time there is no User entity
	CurrentTripID uint   // One-to-One relationship with Vendor
}
