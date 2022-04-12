package main

import (
	"github.com/gofiber/fiber"

	database "./database"

	route "./router"
)

func main() {

	app := fiber.New()

	database.Connection()

	route.SetupRouter(app)

	app.Listen(":8080")

}
