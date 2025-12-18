package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqLiteDb(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
