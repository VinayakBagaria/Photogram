package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/config"
)

type Configuration interface {
	Dsn() string
}

type configuration struct {
	dsn string
}

func NewConfiguration() Configuration {
	var cfg configuration
	dbUser := config.GetEnvString("postgres.user")
	dbPass := config.GetEnvString("postgres.password")
	dbHost := config.GetEnvString("postgres.host")

	dbPort, err := strconv.Atoi(config.GetEnvString("postgres.port"))
	if err != nil {
		log.Fatalln("Unable to load env")
	}

	dbName := config.GetEnvString("postgres.dbname")
	cfg.dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	return &cfg
}

func (c *configuration) Dsn() string {
	return c.dsn
}
