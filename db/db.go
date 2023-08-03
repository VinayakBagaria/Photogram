package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg Configuration) (*gorm.DB, error) {
	dialector := postgres.New(postgres.Config{
		DSN:                  cfg.Dsn(),
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Picture{})

	return db, nil
}
