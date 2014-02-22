package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	db *sqlx.DB
}

type Domain struct {
	Id         int
	DomainId   int `db:"domain_id"`
	Url        string
	Title      string
	Status     string
	PreText    string `db:"pre_text"`
	PostText   string `db:"post_text"`
	Analytics  string
	HeaderText string `db:"header_text"`
}

type Board struct {
	Id int
}

type Post struct {
	Id int
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init(config *Config) error {
	var err error
	m.db, err = sqlx.Open(config.Database, config.Dsn)
	return err
}

func (m *Model) getDomain(host string) *Domain {
	domain := &Domain{}
	err := m.db.Get(domain, "SELECT * FROM edomain WHERE url = $1", host)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return domain
}

func (m *Model) getBoards() *[]Board {
	boards := &[]Board{}
	return boards
}
func (m *Model) getLatestPosts() *[]Post {
	posts := &[]Post{}
	return posts
}
