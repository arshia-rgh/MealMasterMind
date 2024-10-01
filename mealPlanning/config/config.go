package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var DbConfig DBConfig

func InitDBConfig() error {
	err := godotenv.Load("mealPlanning/.env")

	if err != nil {
		return err
	}

	DbConfig = DBConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
	return nil
}
