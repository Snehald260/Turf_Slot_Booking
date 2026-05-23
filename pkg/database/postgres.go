package database

// Postgres connection helper.
// Will provide:
// - Connection pool setup (sql.DB)
// - MaxOpenConns / MaxIdleConns configuration
// - Health check (Ping)
