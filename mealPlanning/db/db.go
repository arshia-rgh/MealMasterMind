package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("mysql", "user:password@/dbname")

}
