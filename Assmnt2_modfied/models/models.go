package models

import "gorm.io/gorm"

type Stud_data struct {
	gorm.Model
	Id         int64  `json:"id" gorm:"type:integer;not null`
	Name       string `json:"name" gorm:"text;not null`
	Class      string `json:"class" gorm:"text;not null`
	Percentage int64  `json:"percentage" gorm:"type:integer;not null"`
	Rank       int64  `json:"rank" gorm:"type:integer;not null"`
}
