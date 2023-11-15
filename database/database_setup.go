package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func DBConnect() {
	conn_str := "user=postgres password=root dbname=crud port=5432 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn_str)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
