package main

import (
	"fmt"
	"idm/inner/database"
	"log"
)

func main() {
	db, err := database.NewConnect(".env")
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
