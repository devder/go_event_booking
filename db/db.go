package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // keep the import when saving file
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // ensure the DB var is the one in the outta scope

	if err != nil {
		panic("failed to connect to the DB")
	}

	DB.SetMaxOpenConns(4)
	DB.SetMaxIdleConns(2)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("failed to create users table")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER
			FOREIGN KEY(userId) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("failed to create events table")
	}
}
