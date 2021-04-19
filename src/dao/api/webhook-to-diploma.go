package api

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/dao/interfaces"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

type WebhookData struct {
	Login               string `json:"login"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	BirthDate           string `json:"birth_date"`
	AlumnizedCursusUser int    `json:"alumnized_cursus_user"`
}

func WebhookToDiploma(body io.ReadCloser) (interfaces.Diploma, error) {
	var diplomaInterface interfaces.Diploma
	var diplomaData diplomas.DiplomaImpl
	var webhookData WebhookData
	err := json.NewDecoder(body).Decode(&webhookData)
	if err != nil {
		tools.LogsError(err)
		return diplomaInterface, err
	}
	level, skills, err := GetCursusUser(webhookData.AlumnizedCursusUser)
	if err != nil {
		tools.LogsError(err)
		return diplomaInterface, err
	}
	diplomaData.FirstName = webhookData.FirstName
	diplomaData.LastName = webhookData.LastName
	diplomaData.BirthDate = webhookData.BirthDate
	diplomaData.AlumniDate = time.Now().Format("2006-01-02")
	diplomaData.Level = level
	diplomaData.Skills = skills
	diplomaData.Counter = 0
	log.WithFields(diplomaData.LogFields()).Debug("Webhook to Diploma success")
	diplomaInterface = diplomaData
	return diplomaInterface, nil
}
