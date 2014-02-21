package main

import (
	"fmt"
)

type GoForIt struct {
	Config *Config
	Model  *Model
}

func NewGoForIt() *GoForIt {
	return &GoForIt{
		Config: NewConfig(),
		Model:  NewModel(),
	}
}

func (gfi *GoForIt) Init(configFile string) error {
	err := gfi.Config.Load(configFile)
	if err != nil {
		fmt.Printf("Cannot load config file! Error: %s\n", err.Error())
		return err
	}
	err = gfi.Model.Init(gfi.Config)
	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return err
	}
	return nil
}
