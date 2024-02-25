package model

type Agent struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"` // At the time there is no User entity
	CurrentOrderID uint   // One-to-One relationship with Vendor
}
