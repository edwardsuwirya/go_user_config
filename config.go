package main

import "fmt"

type Config struct {
	dbHost     string
	dbPort     string
	dbName     string
	dbUser     string
	dbPassword string
}

func GetDataSourceName(c Config) (dsn string) {
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.dbUser, c.dbPassword, c.dbHost, c.dbPort, c.dbName)
	return
}
