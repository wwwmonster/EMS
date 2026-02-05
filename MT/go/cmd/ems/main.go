package main

import (
	"database/sql"
	"ems/mt/golang/emstest"
	"ems/mt/golang/initializers"
	"ems/mt/golang/router"
	"fmt"
	"log"

	"time"
)

func init() {
	initializers.LoadEvnVariables()
	initializers.CreateConnectionPool()
}

func main() {
	// port := os.Getenv("PORT")
	startEms()
	// mainemstest()
}

func mainemstest() {
	emstest.Test()
	fmt.Println("--------")
}

func startEms() {
	fmt.Println("Start EMS")
	emsRouter := router.SetupEMGRouter()
	router.CreateJwtToken("alex")
	emsRouter.Run()

}

// var emsDb *sql.DB = initializers.InitDB() // Global variable to hold the connection pool

func initDB() {
	connStr := "user=pquser dbname=pqgotest sslmode=disable"
	var err error
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
}
