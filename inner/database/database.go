package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

type ConfigError struct {
	message string
}

func (e *ConfigError) Error() string {
	return e.message
}

type DbError struct {
	message string
}

func (e *DbError) Error() string {
	return e.message
}

func NewConnect(path string) (*sqlx.DB, error) {
	cfg, err := load(path)
	if err != nil {
		return nil, &ConfigError{message: fmt.Sprintf("failed to load config from %s", path)}
	}
	cnn, err := sqlx.Connect(cfg.driverName, cfg.dsn)
	if err != nil {
		return nil, &DbError{message: fmt.Sprintf("failed to connect to %s", err.Error())}
	}
	return cnn, nil
}

type dbCfg struct {
	driverName string
	dsn        string
}

func load(path string) (dbCfg, error) {
	err := godotenv.Load(path)
	if err != nil {
		return dbCfg{}, err
	}
	return dbCfg{
		driverName: os.Getenv("DB_DRIVER_NAME"),
		dsn:        os.Getenv("DB_DSN"),
	}, nil
}
