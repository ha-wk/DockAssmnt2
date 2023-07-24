package routes

import (
	"example/Assmnt2_modified/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Listall)
	app.Get("/:id", handlers.GetbyId)

	app.Post("/new", handlers.Createnew)

	app.Put("/:id", handlers.UpdatebyId)

	app.Delete("/:id", handlers.DeletebyId)

}
