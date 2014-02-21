package main

import (
	"fmt"
)

type App struct {
	Config *Config
	Model  *Model
}

func NewApp() *App {
	return &App{
		Config: NewConfig(),
		Model:  NewModel(),
	}
}

func (app *App) Init(configFile string) error {
	err := app.Config.Load(configFile)
	if err != nil {
		fmt.Printf("Cannot load config file! Error: %s\n", err.Error())
		return err
	}
	err = app.Model.Init(app.Config)
	if err != nil {
		fmt.Printf("Cannot open database! Error: %s\n", err.Error())
		return err
	}
	return nil
}
