package handlers

import (
	"example/Assmnt2_modified/database"
	"example/Assmnt2_modified/models"

	"github.com/gofiber/fiber/v2"
)

var dbx database.Dbstruct

func Listall(c *fiber.Ctx) error {
	student_db, err := dbx.ListAll()
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

	err := dbx.CreateNew(newStudent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(newStudent)
}

func GetbyId(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := dbx.GetByID(id)
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

	err := dbx.UpdateByID(id, &newStudent)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	return c.JSON(newStudent)
}

func GetAverageMarks(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := dbx.FetchStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	averageMarks := student.CalculateAverage()

	return c.JSON(fiber.Map{"average_marks": averageMarks})
}

func GetPercentage(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := dbx.FetchStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	stud_percentage := student.CalculatePercentage()

	return c.JSON(fiber.Map{"percentage": stud_percentage})
}

func DeletebyId(c *fiber.Ctx) error {
	id := c.Params("id")

	err := dbx.DeleteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}
