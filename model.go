package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Model struct {
	db *sql.DB
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init(config *Config) error {
	var err error
	m.db, err = sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d slmode=disable",
			config.dbuser,
			config.dbpassword,
			config.dbname,
			config.dbhost,
			config.dbport))
	return err
}

func (m *Model) getDomain(host string) {

}

func (m *Model) getBoards() {
}
func (m *Model) getLatestPosts() {
}
