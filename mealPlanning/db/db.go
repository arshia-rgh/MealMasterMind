package db

import (
	"database/sql"
	"fmt"
	"mealPlanning/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	err := config.InitDBConfig()

	if err != nil {
		panic("could not fetch data from .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DbConfig.DBUser,
		config.DbConfig.DBPassword,
		config.DbConfig.DBHost,
		config.DbConfig.DBPort,
		config.DbConfig.DBName,
	)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("could not open the database connection")
	}

	if err = DB.Ping(); err != nil {
		panic("could not ping the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = applyMigrations()
	if err != nil {
		panic(err)
	}

}
