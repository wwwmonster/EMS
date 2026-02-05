package initializers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *sql.DB {
	connStr := "postgres://admin:123456@localhost:5432/Angular18"
	// var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Configure the connection pool
	db.SetMaxOpenConns(25)                 // Limit total connections
	db.SetMaxIdleConns(25)                 // Keep idle pool full
	db.SetConnMaxLifetime(5 * time.Minute) // Close connections older than 5 minutes

	fmt.Println("Database connection pool established and configured")

	return db
}

func CreateConnectionPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	// Parse the connection string into a Config
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Optional: Configure pool settings
	config.MaxConns = 10 // Maximum number of connections
	config.MinConns = 2  // Minimum number of connections
	// config.MaxConnLifetime = time.Hour // Max time a connection can be open

	// Establish the connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Verify connectivity with a health check
	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("pool health check failed: %w", err)
	}

	return pool, nil
}
