package model

type Agent struct {
	ID              uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name            string `gorm:"not null" json:"name"`
	CurrentReportID uint   `json:"CurrentReportID"` // One-to-One relationship with Vendor
	IsAvailable     bool   `gorm:"default:true" json:"isAvailable"`
}
