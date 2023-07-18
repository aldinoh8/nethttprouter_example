package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeDb(dburl string) *sql.DB {
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
