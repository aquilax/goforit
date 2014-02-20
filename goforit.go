package main

import (
	"fmt"
	"database/sql"
) 

type GoForIt struct {
	Config *Config
	Db *sql.DB
}

func NewGoForIt() (*GoForIt) {
	return &GoForIt{
		Config: NewConfig(),
		Db: nil,
	}
}

func (gfi *GoForIt) Init(configFile string) (error){
	gfi.Config.dbuser = "test"
	gfi.Config.dbpassword = "test"
	gfi.Config.dbname = "test"
	gfi.Config.dbhost = "localhost"
	gfi.Config.dbport = 5432
	var err error
	gfi.Db, err = InitDb(gfi.Config)
	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return err
	}
	return nil
}
