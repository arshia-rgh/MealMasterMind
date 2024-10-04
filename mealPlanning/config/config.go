package config

import (
	"os"

	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
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

func InitDBConfig() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	DBCONFIG = DBConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
	return nil
}

func InitCorsConfig() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	CORSCONFIG = cors.Config{
		AllowAllOrigins:            false,
		AllowOrigins:               nil,
		AllowOriginFunc:            nil,
		AllowOriginWithContextFunc: nil,
		AllowMethods:               nil,
		AllowPrivateNetwork:        false,
		AllowHeaders:               nil,
		AllowCredentials:           false,
		ExposeHeaders:              nil,
		MaxAge:                     0,
		AllowWildcard:              false,
		AllowBrowserExtensions:     false,
		CustomSchemas:              nil,
		AllowWebSockets:            false,
		AllowFiles:                 false,
		OptionsResponseStatusCode:  0,
	}

	return nil
}
