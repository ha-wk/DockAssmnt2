package models

import "gorm.io/gorm"

type MarksCalculator interface {
	CalculateAverage() float64
	CalculatePercentage() float64
}

type Stud_data struct {
	gorm.Model
	Id         int64  `json:"id" gorm:"type:integer;not null"`
	Name       string `json:"name" gorm:"text;not null"`
	Class      string `json:"class" gorm:"text;not null"`
	Percentage int64  `json:"percentage" gorm:"type:integer;not null"`
	Rank       int64  `json:"rank" gorm:"type:integer;not null"`
    Phymarks  int64 `json:"phymarks" gorm:"type:integer;not null"`
	Chemmarks int64 `json:"chemmarks" gorm:"type:integer;not null"`
	Mathmarks int64 `json:"mathmarks" gorm:"type:integer;not null"`
}

func (s *Stud_data) CalculateAverage() float64 {
	totalMarks := s.Phymarks + s.Chemmarks + s.Mathmarks
	return float64(totalMarks) / 3
}

func (s *Stud_data) CalculatePercentage() float64 {
	totalMarks := s.Phymarks + s.Chemmarks + s.Mathmarks
	percentage := (float64(totalMarks) / 300) * 100

	return percentage
}
