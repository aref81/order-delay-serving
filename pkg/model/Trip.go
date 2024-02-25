package model

const (
	TripStatusDelivered = "DELIVERED"
	TripStatusPicked    = "PICKED"
	TripStatusAtVendor  = "AT_VENDOR"
	TripStatusAssigned  = "ASSIGNED"
)

type Trip struct {
	ID      uint   `gorm:"primaryKey"`
	OrderID uint   `gorm:"not null"` // One-to-One relationship with Order
	Status  string `gorm:"not null; type:varchar(10)"`
}
