package database

import (
	"errors"
	"fmt"
	"menchaca-health/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNoRecord is returned when a requested record is not found
var ErrNoRecord = errors.New("record not found")

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	// Initialize the database connection
	err := config.InitDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	// Use the global connection pool
	if config.Conn == nil {
		return nil, errors.New("database connection not initialized")
	}
	return &Database{DB: config.Conn}, nil
}

func (d *Database) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}