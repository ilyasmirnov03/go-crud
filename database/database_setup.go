package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func DBConnect(user string, password string, host string) {
	conn_str := fmt.Sprintf(
		"user=%s password=%s dbname=crud host=%s sslmode=disable",
		user, password, host,
	)
	db, err := sql.Open("postgres", conn_str)
	if err != nil {
		log.Fatal("Error opening SQL connection", err)
	}
	DB = db
}
