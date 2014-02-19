package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var db_ *sql.DB

func GetDb() *sql.DB {
	if db_ != nil {
		return db_
	}
	dbc, err := sql.Open("postgres", "user=%s password=%s dbname=%s host=%s port=%d")

	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return nil
	}
	db_ = dbc
	return db_
}
