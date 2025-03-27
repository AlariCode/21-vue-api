package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./bookmarks.db")
	if err != nil {
		return err
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		alias TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS bookmarks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category_id INTEGER,
		url TEXT NOT NULL,
		title TEXT,
		image TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(createTables)
	return err
}
