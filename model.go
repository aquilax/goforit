package main

import (
	"database/sql"

	//	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	db *sql.DB
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init(config *Config) error {
	var err error
	m.db, err = sql.Open(config.database, config.dsn)
	return err
}

func (m *Model) getDomain(host string) {

}

func (m *Model) getBoards() {
}
func (m *Model) getLatestPosts() {
}
