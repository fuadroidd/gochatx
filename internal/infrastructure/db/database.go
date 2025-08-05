package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil

}
