package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(new(Agent), new(DelayReport), new(Vendor), new(Order), new(Trip))
	if err != nil {
		return err
	}

	return nil
}
