package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// pgx stdlib driver registers itself as "pgx" with database/sql.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Config holds the settings needed to open a Postgres connection pool.
type Config struct {
	DSN             string        // PostgreSQL connection string (DSN)
	MaxOpenConns    int           // max number of open connections to the DB
	MaxIdleConns    int           // max number of idle connections kept in the pool
	ConnMaxLifetime time.Duration // max time a connection may be reused
}

// New opens a Postgres connection pool, verifies connectivity with a Ping,
// and returns the *sql.DB. The caller is responsible for closing it.
func New(ctx context.Context, cfg Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	// Pool tuning. sql.Open does not actually connect; these configure the pool.
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	// Verify the connection is actually reachable.
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		// Close the pool we just opened so we don't leak it on failure.
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}
