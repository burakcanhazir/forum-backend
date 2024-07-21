package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./myproject.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tables if not exists
	createTables()

	log.Println("Database connection established")

	_, err = DB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
	}
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
