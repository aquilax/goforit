package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	database string
	dsn      string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(fileName string) error {
	file, error := os.Open(fileName)
	if error != nil {
		return error
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&c)
}
