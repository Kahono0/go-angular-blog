package db

import (
	"blog/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dbName string) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}

	// migrations
	db.AutoMigrate(
		&models.Blog{},
	)

	DB = db
	return nil
}
