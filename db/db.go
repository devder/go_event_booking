package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // keep the import when saving file
)

var DB sql.DB

func initDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("failed to connect to the DB")
	}

	DB.SetMaxOpenConns(4)
	DB.SetMaxIdleConns(2)
}
