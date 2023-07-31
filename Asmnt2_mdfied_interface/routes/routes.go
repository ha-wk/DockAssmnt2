package routes

import (
	"example/Assmnt2_modified/handlers"
	//"example/Assmnt2_modified/database"

	"github.com/gofiber/fiber/v2"
)  


func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Listall)
	app.Get("/:id", handlers.GetbyId)

	app.Post("/new", handlers.Createnew)

	app.Put("/:id", handlers.UpdatebyId)

	app.Get("/:id/average-marks", handlers.GetAverageMarks)

	app.Get("/:id/percentage", handlers.GetPercentage)

	app.Delete("/:id", handlers.DeletebyId)

}
