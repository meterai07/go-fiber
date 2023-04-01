package main

import (
	"belajar-go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetUser)

	app.Post("/user", handlers.CreateUser)
}
