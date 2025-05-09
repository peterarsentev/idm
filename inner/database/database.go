package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

func NewConnect() (*sqlx.DB, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}
	return sqlx.MustConnect(cfg.driverName, cfg.dsn), nil
}

type dbCfg struct {
	driverName string
	dsn        string
}

func load() (dbCfg, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return dbCfg{}, err
	}
	return dbCfg{
		driverName: os.Getenv("DB_DRIVER_NAME"),
		dsn:        os.Getenv("DB_DSN"),
	}, nil
}
