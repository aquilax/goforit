package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	//	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	STATUS_ENABLED = 1
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

type GroupedBoard struct {
	Name   string
	Boards *[]Board
}

type GroupedBoards []GroupedBoard

type GBIndex map[string]int

type Board struct {
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
	err := m.db.Get(domain,
		"SELECT * FROM edomain"+
			" WHERE url = $1", host)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return domain
}

func (m *Model) getGroupedBoards(domainId int) *GroupedBoards {
	boards := &[]Board{}
	err := m.db.Select(boards,
		"SELECT * FROM forum"+
			" WHERE status = $1"+
			" AND domain_id = $2"+
			" ORDER BY  group_name, sorder", STATUS_ENABLED, domainId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	groupedBoards := &GroupedBoards{}
	gbi := GBIndex{}
	for _, board := range *boards {
		groupedBoards.add(gbi, board)
	}
	return groupedBoards
}

func (gb *GroupedBoards) add(gbi GBIndex, b Board) {
	i, ok := gbi[b.GroupName]
	if !ok {
		*gb = append(*gb, GroupedBoard{
			Name:   b.GroupName,
			Boards: &[]Board{},
		})
		i = len(*gb) - 1
		gbi[b.GroupName] = i
	}
	boards := (*gb)[i].Boards
	(*boards) = append((*boards), b)
}

func (m *Model) getLatestPosts(domainId int, limit int, offset int) *[]Post {
	posts := &[]Post{}
	return posts
}
