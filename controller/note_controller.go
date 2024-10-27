package controller

import (
	"vivo/data/request"
	"vivo/data/response"
	"vivo/helper"
	"vivo/service"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	NoteService service.NoteService
}

func NewNoteController(NoteService service.NoteService) NoteController {
	return NoteController{NoteService: NoteService}
}

func (controller *NoteController) Create(c *fiber.Ctx) error {
	var createNoteRequest request.CreateNoteRequest
	err := c.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.NoteService.Create(createNoteRequest)

	webRespoponse := response.Response{
		Code:   200,
		Status: "OK",
		Message: "Create Note Success",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(webRespoponse)
}

func (controller *NoteController) Update(c *fiber.Ctx) error {
	var updateNoteRequest request.UpdateNoteRequest
	err := c.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)

	controller.NoteService.Update(updateNoteRequest)

	webRespoponse := response.Response{
		Code:   200,
		Status: "OK",
		Message: "Update Note Success",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(webRespoponse)
}

func (controller *NoteController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	helper.ErrorPanic(err)

	controller.NoteService.Delete(id)

	webRespoponse := response.Response{
		Code:   200,
		Status: "OK",
		Message: "Delete Note Success",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(webRespoponse)
}

func (controller *NoteController) FindAll(c *fiber.Ctx) error {
	noteResponses := controller.NoteService.FindAll()
	webRespoponse := response.Response{
		Code:   200,
		Status: "OK",
		Message: "Get All Note Success",
		Data:   noteResponses,
	}

	return c.Status(fiber.StatusOK).JSON(webRespoponse)
}

func (controller *NoteController) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	helper.ErrorPanic(err)

	noteResponse := controller.NoteService.FindByID(id)
	webRespoponse := response.Response{
		Code:   200,
		Status: "OK",
		Message: "Get Note Success",
		Data:   noteResponse,
	}

	return c.Status(fiber.StatusOK).JSON(webRespoponse)
}