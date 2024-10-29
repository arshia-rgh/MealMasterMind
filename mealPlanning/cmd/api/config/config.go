package config

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var CORSCONFIG cors.Config
var DBCONFIG DBConfig

func InitConfigs() error {

	DBCONFIG = DBConfig{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBName:     os.Getenv("POSTGRES_DB"),
	}

	CORSCONFIG = cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return nil
}
