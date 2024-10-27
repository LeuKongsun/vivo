package main

import (
	"fmt"
	"log"
	"vivo/controller"
	"vivo/initializer"
	"vivo/model"
	"vivo/repository"
	"vivo/router"
	"vivo/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Print("Server is running...")
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	db := initializer.ConnectDB(&config)
	validate := validator.New()
	db.Table("notes").AutoMigrate(&model.Note{})

	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteServiceImpl(noteRepository, validate)
	noteController := controller.NewNoteController(noteService)
	routes := router.NewRouter(&noteController)

	app := fiber.New()
	
	app.Mount("/api", routes)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	log.Fatal(app.Listen(":8000"))
}
