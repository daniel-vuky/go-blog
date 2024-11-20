package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

var connectionPool *pgxpool.Pool
var databaseOnce sync.Once

// ConnectToPgxPool
// Connect to the database pool
// @return *pgxpool.Pool
func (config *Config) ConnectToPgxPool() (*pgxpool.Pool, error) {
	var err error
	databaseOnce.Do(func() {
		connectionPool, err = pgxpool.New(context.Background(), config.GetDatabaseSource())
	})

	return connectionPool, err
}
