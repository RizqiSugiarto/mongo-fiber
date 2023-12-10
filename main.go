package main

import (
	"golang-mongo-fiber/configs"
	"golang-mongo-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":8080")
}
