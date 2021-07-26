package storage

import (
	"database/sql"

	_ "github.com/ClickHouse/clickhouse-go"
)

// Storage storage instance
type Storage struct {
	Client *sql.DB
}

// NewStorage returns a new instance
func NewStorage(dbURI string) (*Storage, error) {
	connect, err := sql.Open("clickhouse", dbURI)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: connect,
	}, nil
}