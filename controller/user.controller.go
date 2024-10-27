package controller

import (
	"vivo/model"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
