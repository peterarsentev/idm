package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"idm/inner/database"
	"log"
)

func main() {
	db, err := database.NewConnect()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	fmt.Println("Database connection established")
}
