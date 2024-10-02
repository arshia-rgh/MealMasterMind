package db

import (
	"os"
	"path/filepath"
)

func applyMigrations() error {
	migrationsDir := "../migrations"

	files, err := os.ReadDir(migrationsDir)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		migrations, err := os.ReadFile(filePath)

		if err != nil {
			return err
		}

		_, err = DB.Exec(string(migrations))

		if err != nil {
			return err
		}
	}

	return nil

}
