package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDb(config *Config) (*sql.DB, error) {
	return sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d slmode=disable",
			config.dbuser,
			config.dbpassword,
			config.dbname,
			config.dbhost,
			config.dbport))
}
