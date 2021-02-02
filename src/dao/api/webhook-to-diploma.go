package api

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"io"
	"log"
	"time"
)

type WebhookData struct {
	Login				string	`json:"login"`
	FirstName			string	`json:"first_name"`
	LastName			string	`json:"last_name"`
	AlumnizedCursusUser	int		`json:"alumnized_cursus_user"`
}

func WebhookToDiploma(body io.ReadCloser) (diplomas.Diploma, error) {
	var newDiploma diplomas.Diploma
	var webhookData WebhookData
	err := json.NewDecoder(body).Decode(&webhookData)
	log.Println("in webhook->dp:", err)
	if err != nil {
		return newDiploma, err
	}
	level, skills, err := GetCursusUser(webhookData.AlumnizedCursusUser)
	log.Println("in webhook->dp:", err)
	if err != nil {
		return newDiploma, err
	}
	birthDate, err := GetBirthdate(webhookData.Login)
	log.Println("in webhook->dp:", err)
	if err != nil {
		return newDiploma, err
	}
	newDiploma.FirstName = webhookData.FirstName
	newDiploma.LastName = webhookData.LastName
	newDiploma.BirthDate = birthDate
	newDiploma.AlumniDate = time.Now().Format("2006-01-02")
	newDiploma.Level = level
	newDiploma.Skills = skills
	log.Println(newDiploma.String(), level, skills)
	return newDiploma, nil
}
