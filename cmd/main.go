package main

import (
	"belajar-go-fiber/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
