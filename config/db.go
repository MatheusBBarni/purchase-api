package config

import (
	"github.com/matheusBBarni/purchase-api/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("purchase-api.db"), &gorm.Config{})

	db.AutoMigrate(&domain.Purchase{})

	return db, err
}
