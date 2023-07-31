package main

import (
	"example/Assmnt2_modified/database"
	"example/Assmnt2_modified/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// str, _ := os.LookupEnv("DB_HOST")
	str := os.Getenv("DB_HOST")
	fmt.Printf("str: %v\n", str)

	database.ConnectDb()


	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
