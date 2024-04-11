package media

import (
	"github.com/gofiber/fiber/v3"

	"github.com/khelechy/memorize/services"
)

func FetchMedia(c fiber.Ctx) error {
	userId := c.Params("userId")

	if len(userId) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User id is required"})
	}

	mediaLocation, err := services.FetchUserUploads(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	data := map[string]interface{}{
		"mediaLocation": mediaLocation,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Media successfully retreived", "data": data})
}

func UploadMedia(c fiber.Ctx) error {
	userId:= c.Params("userId")

	if len(userId) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User id is required"})
	}
	
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = services.HandleMediaUpload(userId, file, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("File uploaded")
}
