package datasources

import (
	"OrderDelayServing/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func InitPostgres(conf *config.Config) (*gorm.DB, error) {
	port := strconv.Itoa(conf.Database.Port)
	dsn := "host=" + conf.Database.Addr + " user=" + conf.Database.User + " password=" + conf.Database.Password + " dbname=" + conf.Database.DBName + " port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
