package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	log.Println("Successfully connecte to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&Picture{})

	return db, nil
}
