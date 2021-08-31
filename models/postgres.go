package models

import (
	"database/sql"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=yasmins dbname=go_tutorial sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	return db
}
