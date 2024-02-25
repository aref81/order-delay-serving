package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(new(Agent), new(DelayReport), new(Order), new(Trip), new(Vendor))
	if err != nil {
		return err
	}

	return nil
}
