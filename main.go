package main

import (
	"fiber-go-api/configs"
	"fiber-go-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello world"})
	// })
	// connect database

	configs.ConnectDB()

	//routes

	routes.UserRoute(app)

	app.Listen(":6000")
}
