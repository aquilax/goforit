package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)


func InitDb(config *Config) (*sql.DB, error) {
	return sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
			config.dbuser,
			config.dbpassword,
			config.dbname,
			config.dbhost,
			config.dbport))
}

