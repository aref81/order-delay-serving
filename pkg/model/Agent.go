package model

type Agent struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	CurrentOrderID uint   // One-to-One relationship with Vendor
}
