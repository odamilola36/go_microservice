package dbCon

import (
	"fmt"
	"os"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")
	cfg.dsn = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbName)
	return &cfg
}

func (c config) Dsn() string {
	return c.dsn
}

func (c config) DbName() string {
	return c.dbName
}
