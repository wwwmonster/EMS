package initializers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
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

var emsPgxConnPool *pgxpool.Pool

func CreateConnectionPool() {
	// Parse the connection string into a Config
	var ctx = context.Background()
	config, err := pgxpool.ParseConfig("postgres://admin:123456@localhost:5432/Angular18")
	// config, err := pgxpool.ParseConfig(os.Getenv("PG_CONN_STRING"))
	if err != nil {
		log.Fatal("unable to parse connection string: %w", err)
		os.Exit(1)
	}

	// Optional: Configure pool settings
	config.MaxConns = 10 // Maximum number of connections
	config.MinConns = 2  // Minimum number of connections
	// config.MaxConnLifetime = time.Hour // Max time a connection can be open

	// Establish the connection pool
	emsPgxConnPool, err = pgxpool.NewWithConfig(ctx, config)

	fmt.Println("-----------ppp-------------")
	if err != nil {
		log.Fatal("unable to create connection pool: %w", err)
		os.Exit(1)
	}

	// Verify connectivity with a health check
	err = emsPgxConnPool.Ping(ctx)
	if err != nil {
		emsPgxConnPool.Close()
		log.Fatal("pool health check failed: %w", err)
		os.Exit(1)
	}
}

func GetEmsPgxConnPool() *pgxpool.Pool {
	return emsPgxConnPool
}
