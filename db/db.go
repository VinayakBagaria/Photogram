package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg Configuration) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Dsn()))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Picture{})

	return db, nil
}
