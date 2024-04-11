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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully", "data": data})
}

func GetUserQr(c fiber.Ctx) error {
	userId := c.Params("userId")

	if len(userId) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User id is required"})
	}

	qr, ss, err := services.SaveUserQr(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	data := map[string]interface{}{
		"qrUrl": qr,
		"siteUrl": ss,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User QR generated successfully", "data": data})
}
