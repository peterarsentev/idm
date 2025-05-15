package repositories

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"idm/inner/models"
	repository "idm/inner/repositories"
	"log"
	"testing"
)

func TestCreateFilter(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	if clearErr := clearDb(db); clearErr != nil {
		log.Printf("Error while truncating tables: %v", clearErr)
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewFilterRepository(db)
	filter := models.Filter{
		Name: "Test Filter",
	}
	ID, err := repo.Add(context.Background(), filter)
	a.NoError(err, "Failed to add filter")
	createdFilter, err := repo.FindByID(context.Background(), ID)
	a.NoError(err, "Failed to retrieve filter")
	a.Equal(filter.Name, createdFilter.Name, "Filter names should match")
}

func TestFindAllFilters(t *testing.T) {
	a := assert.New(t)
	db, err := database.NewDb(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	repo := repository.NewFilterRepository(db)
	filters, err := repo.FindAll(context.Background())
	a.NoError(err, "Failed to fetch filters")
	a.NotEmpty(filters, "Filters should not be empty")
}
