package db

import (
	"database/sql"
	"embed"
	"sort"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunMigrations(db *sql.DB) error {
	rows, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return err
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Name() < rows[j].Name()
	})

	for _, file := range rows {
		path := "migrations/" + file.Name()
		content, err := migrationFiles.ReadFile(path)
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(content)); err != nil {
			return err
		}
	}

	return nil
}
