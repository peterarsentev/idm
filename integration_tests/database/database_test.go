package database

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"log"
	"testing"
	"time"
)

func TestWhenNotFoundEnv(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewConnect(fmt.Sprintf(".env.%d", time.Nanosecond))
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	a.NotNil(err)

	// Check if the error is of type ConfigError
	var configErr *database.ConfigError
	a.True(errors.As(err, &configErr), "Expected ConfigError, got %T", err)
}

func TestWhenEmptyEnv(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewConnect(".empty")
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	a.NotNil(err)
	var configErr *database.DbError
	a.True(errors.As(err, &configErr), "Expected DbError, got %T", err)
}

func TestWhenConnectionOk(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewConnect(".env.ok")
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	a.Nil(err)
}
