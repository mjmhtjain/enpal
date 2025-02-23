package client

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var postgresDB *sql.DB = nil

func NewDBClient(config *DatabaseConfig) (*sql.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	db, err := sql.Open("postgres", config.GetDSN())
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	postgresDB = db

	return postgresDB, nil
}
