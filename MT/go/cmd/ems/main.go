package main

import (
	"database/sql"
	"ems/mt/golang/router"
	"fmt"
	"log"
	"time"
)

func main() {
	startEms()
}

func startEms() {
	fmt.Println("Start EMS")
	emsRouter := router.SetupEMGRouter()
	router.CreateJwtToken("alex")
	emsRouter.Run(":9090")
}

var db *sql.DB // Global variable to hold the connection pool

func initDB() {
	connStr := "user=pquser dbname=pqgotest sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
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
}
