package main

import (
	"idm/inner/database"
	"idm/inner/handlers"
	"idm/inner/repositories"
	"idm/inner/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	db, err := database.NewConnect(".env")
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	filterHandler := handlers.NewFilterHandler(
		services.NewFilterService(
			repositories.NewFilterRepository(db),
			repositories.NewKeyRepository(db),
			repositories.NewKeyValueRepository(db),
		),
	)
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", handlers.IndexHandler)
	app.Get("/index", handlers.IndexHandler)
	app.Get("/filters/", filterHandler.Filters)
	app.Get("/filter/create", filterHandler.CreateView)
	app.Post("/filter/create", filterHandler.Save)
	app.Get("/filter/view/:id", filterHandler.View)
	log.Fatal(app.Listen(":3000"))
}
