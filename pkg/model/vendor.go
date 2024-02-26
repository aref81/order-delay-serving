package model

type Vendor struct {
	ID     uint    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name   string  `json:"name"`
	Orders []Order `json:"orders"` // One-to-Many relationship with Order
}
