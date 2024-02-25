package model

type Vendor struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Orders []Order // One-to-Many relationship with Order
}
