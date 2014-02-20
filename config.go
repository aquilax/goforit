package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	dbuser     string
	dbpassword string
	dbname     string
	dbhost     string
	dbport     int
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
