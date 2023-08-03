package db

import (
	"fmt"

	"github.com/VinayakBagaria/go-cat-pictures/config"
)

type Configuration interface {
	Dsn() string
}

type configuration struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func NewConfiguration() Configuration {
	var cfg configuration
	cfg.dbUser = config.GetEnvString("postgres.user")
	cfg.dbPass = config.GetEnvString("postgres.password")
	cfg.dbHost = config.GetEnvString("postgres.host")
	cfg.dbPort = config.GetEnvString("postgres.port")
	cfg.dbName = config.GetEnvString("postgres.dbname")
	return cfg
}

func (c configuration) Dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.dbHost, c.dbPort, c.dbUser, c.dbPass, c.dbName)
}
