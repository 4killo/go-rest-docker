package service

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "your my sql url")
	if err != nil {
		log.Fatal("Could not connect to database", err.Error())
	}
	return db
}
