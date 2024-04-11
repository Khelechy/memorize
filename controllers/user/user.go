package user

import (
	"strings"

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
