package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect establishes the connection to the PostgreSQL database
func Connect() {
	var err error
	// Replace the following with your actual PostgreSQL connection string
	connStr := "postgres://postgres:admin@localhost:5432/license?sslmode=disable&search_path=public"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Check if the connection is successful
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	fmt.Println("Connected to the database")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return DB
}

// Close closes the database connection
func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error closing database connection: ", err)
	}
	fmt.Println("Connection closed")
}
