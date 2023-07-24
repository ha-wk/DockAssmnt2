package handlers

import (
	"example/Assmnt2_modified/database"
	"example/Assmnt2_modified/models"

	//"fmt"

	"github.com/gofiber/fiber/v2"
)

func Listall(c *fiber.Ctx) error {
	student_db := []models.Stud_data{}
	database.DB.Db.Find(&student_db)

	//var students models.Stud_data
	tempdata := new(models.Stud_data)
	if err := c.BodyParser(tempdata); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&tempdata)

	return c.Status(200).JSON(student_db)
}

func Createnew(c *fiber.Ctx) error {
	newStudent := new(models.Stud_data)
	if err := c.BodyParser(newStudent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&newStudent)

	return c.Status(200).JSON(newStudent)
}

func GetbyId(c *fiber.Ctx) error {
	id := c.Params("id")

	var student models.Stud_data
	result := database.DB.Db.First(&student, id)
	if result.Error != nil {
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

	var student models.Stud_data
	result := database.DB.Db.First(&student, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	student.Name = newStudent.Name
	student.Class = newStudent.Class
	student.Percentage = newStudent.Percentage
	student.Rank = newStudent.Rank

	database.DB.Db.Save(&student)

	return c.JSON(student)
}

func DeletebyId(c *fiber.Ctx) error {
	id := c.Params("id")

	var student models.Stud_data
	result := database.DB.Db.First(&student, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Student not found"})
	}

	database.DB.Db.Unscoped().Delete(&student)

	return c.Status(fiber.StatusNoContent).SendString("")

}
