package controller

import (
	"vivo/model"
	"vivo/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
    UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
    return UserController{UserService: UserService}
}

func (uc *UserController) GetMe(c *fiber.Ctx) error {
    user := c.Locals("user").(model.UserResponse)

    dbUser, err := uc.UserService.GetUserByID(user.ID.String())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "User not found"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": "success",
        "data": fiber.Map{"user": dbUser},
    })
}
