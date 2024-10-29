package db

import (
	"database/sql"
	"fmt"
	"mealPlanning/cmd/api/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBCONFIG.DBUser,
		config.DBCONFIG.DBPassword,
		config.DBCONFIG.DBHost,
		config.DBCONFIG.DBPort,
		config.DBCONFIG.DBName,
	)
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

}
