package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func createConnection() *sql.DB {
	dsn := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
	con, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return con
}

func DBConnection() *sql.DB {
	if db != nil {
		return db
	}
	return createConnection()
}
