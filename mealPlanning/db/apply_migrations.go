package db

import (
	"os"
	"path/filepath"
	"strings"
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

		strMigrations := string(migrations)

		migrationsSlice := strings.Split(strMigrations, ";")

		for _, sql := range migrationsSlice {
			_, err = DB.Exec(sql)

			if err != nil {
				return err
			}

		}

	}

	return nil

}
