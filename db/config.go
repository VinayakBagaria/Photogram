package db

import (
	"log"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/config"
)

type Configuration interface {
	Dsn() string
	DbName() string
}

type configuration struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Configuration {
	var cfg configuration
	cfg.dbUser = config.GetEnvString("postgres.user")
	cfg.dbPass = config.GetEnvString("postgres.password")
	cfg.dbHost = config.GetEnvString("postgres.host")

	var err error
	cfg.dbPort, err = strconv.Atoi(config.GetEnvString("postgres.port"))
	if err != nil {
		log.Fatalln("Unable to load env")
	}

	cfg.dbName = config.GetEnvString("postgres.dbname")
	cfg.dsn = ""

	return &cfg
}

func (c *configuration) Dsn() string {
	return c.dsn
}

func (c *configuration) DbName() string {
	return c.dbName
}
