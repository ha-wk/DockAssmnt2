package database

import (
	"example/Assmnt2_modified/models"
	"fmt"
	//"gorm.io/gorm"
)

type Dbstruct struct {
}

type DbOperations interface {
	ListAll() ([]models.Stud_data, error)
	CreateNew(newStudent *models.Stud_data) error
	GetByID(id string) (models.Stud_data, error)
	UpdateByID(id string, newStudent *models.Stud_data) error
	FetchStudentByID(id string) (*models.Stud_data, error)
	DeleteByID(id string) error
}

// func New(){

// }

func (d Dbstruct) ListAll() ([]models.Stud_data, error) {
	var students []models.Stud_data
	err := DB.Db.Find(&students).Error
	return students, err
}

func (d Dbstruct) CreateNew(newStudent *models.Stud_data) error {
	return DB.Db.Create(newStudent).Error
}

func (d Dbstruct) GetByID(id string) (models.Stud_data, error) {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	return student, result.Error
}

func (d Dbstruct) UpdateByID(id string, newStudent *models.Stud_data) error {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	if result.Error != nil {
		return result.Error
	}

	student.Name = newStudent.Name
	student.Class = newStudent.Class
	student.Percentage = newStudent.Percentage
	student.Rank = newStudent.Rank
	student.Phymarks = newStudent.Phymarks
	student.Chemmarks = newStudent.Chemmarks
	student.Mathmarks = newStudent.Mathmarks

	return DB.Db.Save(&student).Error
}

func (d Dbstruct) FetchStudentByID(id string) (*models.Stud_data, error) {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	if result.Error != nil {
		return nil, fmt.Errorf("student not found")
	}
	return &student, nil
}

func (d Dbstruct) DeleteByID(id string) error {
	var student models.Stud_data
	result := DB.Db.First(&student, id)
	if result.Error != nil {
		return result.Error
	}

	return DB.Db.Unscoped().Delete(&student).Error
}
