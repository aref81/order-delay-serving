package model

type Agent struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;"json:"id"`
	Name           string `gorm:"not null"json:"name"`
	CurrentOrderID uint   `json:"CurrentOrderID"json:"currentOrderID"` // One-to-One relationship with Vendor
}
