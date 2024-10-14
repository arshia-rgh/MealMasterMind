package db

import (
	"database/sql"
	"fmt"
	"mealPlanning/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBCONFIG.DBUser,
		config.DBCONFIG.DBPassword,
		config.DBCONFIG.DBHost,
		config.DBCONFIG.DBPort,
		config.DBCONFIG.DBName,
	)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

}
