package main

type Config struct {
	dbuser string
	dbpassword string
	dbname string
	dbhost string
	dbport int
}

func NewConfig() (*Config) {
	return &Config{}
}

