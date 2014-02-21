package main

import (
	"github.com/jmoiron/sqlx"
	//	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	db *sqlx.DB
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init(config *Config) error {
	var err error
	m.db, err = sqlx.Open(config.Database, config.Dsn)
	return err
}

func (m *Model) getDomain(host string) {

}

func (m *Model) getBoards() {
}
func (m *Model) getLatestPosts() {
}
