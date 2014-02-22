package main

import (
	"fmt"
	"time"
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
	Id        int
	DomainId  int `db:"domain_id"`
	Title     string
	Body      string
	Status    int
	Sorder    int
	Created   time.Time
	Oid       int
	Updated   time.Time
	Topics    int
	GroupName string `db:"group_name"`
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
	err := m.db.Get(domain,
		"SELECT * FROM edomain"+
			" WHERE url = $1", host)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return domain
}

func (m *Model) getBoards(domainId int) *[]Board {
	boards := &[]Board{}
	err := m.db.Get(boards,
		"SELECT * FROM forum"+
			" WHERE status = 1"+
			" AND domain_id = $1"+
			" ORDER BY  group_name, sorder", domainId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return boards
}
func (m *Model) getLatestPosts(domainId int, limit int, offset int) *[]Post {
	posts := &[]Post{}
	return posts
}
