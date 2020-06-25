package models

import (
	"log"
	"time"
)

type Diploma struct {
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	BirthDate	time.Time	`json:"birth_date"`
	AlumniDate	time.Time	`json:"alumni_date"`
	Level		float64		`json:"level"`
	Skills		[]float64	`json:"skills"`
}

func NewDiploma(new Diploma) {
	log.Print("Enter in NewDiploma")
	log.Println("First Name:", new.FirstName)
	log.Println("Last Name:", new.LastName)
	log.Println("Birth Date:", new.BirthDate)
	log.Println("Alumni Date:", new.AlumniDate)
	log.Println("Level:", new.Level)
	log.Println("Skills:", new.Skills)
}
