package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg Configuration) (*gorm.DB, error) {
	fmt.Println(cfg.Dsn())
	db, err := gorm.Open(postgres.Open(cfg.Dsn()))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Picture{})

	return db, nil
}
