package main

import (
	"database/sql"
	"fmt"
)

type GoForIt struct {
	Config *Config
	Db     *sql.DB
}

func NewGoForIt() *GoForIt {
	return &GoForIt{
		Config: NewConfig(),
		Db:     nil,
	}
}

func (gfi *GoForIt) Init(configFile string) error {
	err := gfi.Config.Load(configFile)
	if err != nil {
		fmt.Printf("Cannot load config file! Error: %s\n", err.Error())
		return err
	}
	gfi.Db, err = InitDb(gfi.Config)
	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return err
	}
	return nil
}
