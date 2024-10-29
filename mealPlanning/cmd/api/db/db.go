package db

import (
	"database/sql"
	"fmt"
	"log"
	"mealPlanning/cmd/api/config"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB() {
	var counts int

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBCONFIG.DBUser,
		config.DBCONFIG.DBPassword,
		config.DBCONFIG.DBHost,
		config.DBCONFIG.DBPort,
		config.DBCONFIG.DBName,
	)
	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres !")
			DB = connection
			return
		}

		if counts > 10 {
			log.Println(err)
			return
		}

		log.Println("Try again in two seconds ...")
		time.Sleep(2 * time.Second)

	}
}
