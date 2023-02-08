package models

// StudentInfo represents the model of a student
type StudentInfo struct {
	ID     string `json:"ID" example:"1"`
	Name   string `json:"Name" example:"Himanshu"`
	Course string `json:"Course" example:"Computer Engineering"`
}
