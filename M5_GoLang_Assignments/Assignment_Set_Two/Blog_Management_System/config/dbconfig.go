package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Initialize the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create blogs table if it does not exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS blogs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			author TEXT NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
