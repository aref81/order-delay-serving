package model

const (
	TripStatusDelivered = "DELIVERED"
	TripStatusPicked    = "PICKED"
	TripStatusAtVendor  = "AT_VENDOR"
	TripStatusAssigned  = "ASSIGNED"
	TripStatusQueued    = "QUEUED"
	TripStatusOnReview  = "ONREVIEW"
	TripStatusReviewed  = "REVIEWED"
)

type Trip struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	OrderID uint   `gorm:"not null" json:"orderID"` // One-to-One relationship with Order
	Status  string `gorm:"not null; type:varchar(10)" json:"status"`
}
