package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewConnection(cfg Configuration) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Dsn())
	if err != nil {
		return nil, err
	}

	return db, nil
}
