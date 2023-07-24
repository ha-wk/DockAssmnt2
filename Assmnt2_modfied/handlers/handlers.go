package handlers

import (
	"example/Assmnt2_modified/database"
	"example/Assmnt2_modified/models"

	"github.com/gofiber/fiber/v2"
)

func Listall(c *fiber.Ctx) error {
	student_db, err := database.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(student_db)
}

func Createnew(c *fiber.Ctx) error {
	newStudent := new(models.Stud_data)
	if err := c.BodyParser(newStudent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := database.CreateNew(newStudent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(newStudent)
}

func GetbyId(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := database.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	return c.JSON(student)
}

func UpdatebyId(c *fiber.Ctx) error {
	id := c.Params("id")

	var newStudent models.Stud_data
	if err := c.BodyParser(&newStudent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	err := database.UpdateByID(id, &newStudent)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	return c.JSON(newStudent)
}

func DeletebyId(c *fiber.Ctx) error {
	id := c.Params("id")

	err := database.DeleteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}
