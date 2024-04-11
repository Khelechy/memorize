package user

import (
	"strings"

	"github.com/khelechy/memorize/services"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func CreateUser(c fiber.Ctx) error {

	uniqueId := uuid.New()
	userId := strings.Replace(uniqueId.String(), "-", "", -1)

	data := map[string]interface{}{
		"userId": userId,
	}
	return c.JSON(fiber.Map{"status": 201, "message": "User created successfully", "data": data})
}

func GetUserQr(c fiber.Ctx) error {
	userId := c.Params("userId")

	qr, err := services.SaveUserQr(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	data := map[string]interface{}{
		"qrUrl": qr,
	}
	
	return c.JSON(fiber.Map{"status": 201, "message": "User QR generated successfully", "data": data})
}
