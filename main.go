package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joshblades/goshorty/inits"
	"github.com/joshblades/goshorty/routes"
	"github.com/joshblades/goshorty/utils"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shortcode", routes.ParseURL)

	// Group mappings
	api := app.Group("/api")
	link := api.Group("/link")

	// Links
	link.Get("/:shortcode", routes.GetLink)
	link.Post("/", routes.CreateLink)
	link.Get("/", routes.GetLinks)
	link.Delete("/:shortcode", routes.DeleteLink)
	link.Patch("/:shortcode", routes.ToggleLinkStatus)
}

func init() {
	utils.LoadEnv()
	inits.ConnectDatabase()
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	listenPort := utils.Getenv("HTTP_LISTEN_PORT", ":3000")

	log.Fatal(app.Listen(listenPort))
}
