package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/khelechy/memorize/controllers"
)

func main() {
	app := fiber.New()

	app.Static("/media/uploads", "/uploads")
	app.Static("/memory", "/views")

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	controllers.RegisterRoutes(app)

	app.Listen(":3000")
}
