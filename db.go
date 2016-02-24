package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDb() (db *sql.DB) {
	db, err := sql.Open("postgres", "password=goacc user=goacc dbname=goacc sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateDb() {
	db := GetDb()
	sqlCreate := `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		balance NUMERIC NOT NULL
	)
	`
	if _, err := db.Exec(sqlCreate); err != nil {
		log.Fatal(err)
	}
}
