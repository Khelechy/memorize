package controllers

import (
	"github.com/gofiber/fiber/v3"

	media "github.com/khelechy/memorize/controllers/media"
	user "github.com/khelechy/memorize/controllers/user"
)

func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api")

	routes.Post("/media/:userId/upload", media.UploadMedia)
	routes.Get("/media/:userId", media.FetchMedia)

	routes.Post("/user", user.CreateUser)
	routes.Get("/user/:userId/qr", user.GetUserQr)

}
