package client

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB = nil

func NewDBClient(config *DatabaseConfig) (*gorm.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	var err error
	postgresDB, err = gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test the connection
	sqlDB, err := postgresDB.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return postgresDB, nil
}
