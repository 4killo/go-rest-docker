package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func connect() *sql.DB {
	fmt.Println("MySql host = " + os.Getenv("DATABASE_HOST"))
	db, err := sql.Open("mysql", "all:123@tcp("+os.Getenv("DATABASE_HOST")+":3306)/customers")
	if err != nil {
		log.Fatal("Could not connect to database", err.Error())
	}
	return db
}
