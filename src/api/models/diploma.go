package models

import "time"

type Diploma struct {
	firstName   string `json:"first_name"`
	lastName   string `json:"last_name"`
	birthDate  time.Time `json:"birth_date"`
	alumniDate time.Time `json:"alumni_date"`
	level       float32 `json:"level"`
	skills      []float32 `json:"skills"`
}