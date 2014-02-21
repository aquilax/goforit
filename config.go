package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database string
	Dsn      string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&c)
}
