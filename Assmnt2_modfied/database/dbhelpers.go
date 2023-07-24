package database

import (
	"example/Assmnt2_modified/models"
	//"gorm.io/gorm"
)

func ListAll() ([]models.Stud_data, error) {
	var students []models.Stud_data
	err := DB.Db.Find(&students).Error
	return students, err
}

func CreateNew(newStudent *models.Stud_data) error {
	return DB.Db.Create(newStudent).Error
}

func GetByID(id string) (models.Stud_data, error) {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	return student, result.Error
}

func UpdateByID(id string, newStudent *models.Stud_data) error {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	if result.Error != nil {
		return result.Error
	}

	student.Name = newStudent.Name
	student.Class = newStudent.Class
	student.Percentage = newStudent.Percentage
	student.Rank = newStudent.Rank

	return DB.Db.Save(&student).Error
}

func DeleteByID(id string) error {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	if result.Error != nil {
		return result.Error
	}

	return DB.Db.Unscoped().Delete(&student).Error
}
