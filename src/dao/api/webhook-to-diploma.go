package api

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/tools"
	"io"
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
	if err != nil {
		tools.LogsError(err)
		return newDiploma, err
	}
	level, skills, err := GetCursusUser(webhookData.AlumnizedCursusUser)
	if err != nil {
		tools.LogsError(err)
		return newDiploma, err
	}
	birthDate, err := GetBirthdate(webhookData.Login)
	if err != nil {
		tools.LogsError(err)
		return newDiploma, err
	}
	newDiploma.FirstName = webhookData.FirstName
	newDiploma.LastName = webhookData.LastName
	newDiploma.BirthDate = birthDate
	newDiploma.AlumniDate = time.Now().Format("2006-01-02")
	newDiploma.Level = level
	newDiploma.Skills = skills
	tools.LogsDev("Webhook -> Diploma: ok")
	return newDiploma, nil
}
