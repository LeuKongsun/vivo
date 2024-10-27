package router

import (
	"fmt"
	"vivo/controller"
	"vivo/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Route("/notes", func(r fiber.Router) {
		r.Use(middleware.DeserializeUser)
		r.Post("/", noteController.Create)
		r.Get("", noteController.FindAll)
	})

	router.Route("/notes/:id", func(r fiber.Router) {
		r.Use(middleware.DeserializeUser)
		r.Get("", noteController.FindByID)
		r.Patch("", noteController.Update)
		r.Delete("", noteController.Delete)
	})

	router.Route("/auth", func(r fiber.Router) {
		r.Post("/register", controller.SignUpUser)
		r.Post("/login", controller.SignInUser)
		r.Get("/logout", middleware.DeserializeUser, controller.LogoutUser)
	})

	router.Get("/users/me", middleware.DeserializeUser, controller.GetMe)

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Vivo project running",
		})
	})

	router.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exist on this server", c.Path()),
		})
	})

	return router
}
