package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"link-share/models"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Link{},
		&models.Collection{},
	)
	if err != nil {
		return
	}
	DB = db
}
