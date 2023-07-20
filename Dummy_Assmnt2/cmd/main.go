package main

import (
	"example/Dummy_Assmnt2/database"
	"example/Dummy_Assmnt2/routes"
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
