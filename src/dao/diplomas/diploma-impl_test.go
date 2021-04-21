package diplomas

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
)

type MockDiplomaImpl struct {
	Id         uuid.UUID `bson:"_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	AlumniDate string    `json:"alumni_date"`
	Level      float64   `json:"level"`
	Skills     []api.Skill   `json:"skills"`
	Counter    int       `json:"counter"`
}

func (dp MockDiplomaImpl) ReadWebhook(body io.ReadCloser) (MockDiplomaImpl, error) {
	dp.FirstName = "Louise"
	dp.LastName = "Pieri"
	dp.BirthDate = "1998-12-27"
	dp.AlumniDate = "2021-01-01"
	dp.Level = 21
	dp.Skills = []api.Skill{{"Security",16.42}, {"Unix",13.87},
		{"Adaptation & creativity",12.7},
		{"Company experience",11.22},
		{"Algorithms & AI",10.38}, {"Group & interpersonal",10.13},
		{"Graphics",7.49}, {"Rigor",6.6},
		{"Imperative programming",5.34},
		{"Technology integration",5.26},
		{"Web",5.2}, {"Organization",5.04},
		{"Network & system administration",4.5},
		{"DB & Data",4.28}, {"Object-oriented programming",4.2}}
	return dp, nil
}

func (dp MockDiplomaImpl) ReadJson(body io.ReadCloser) (MockDiplomaImpl, error) {
	dp.FirstName = "Louise"
	dp.LastName = "Pieri"
	dp.BirthDate = "1998-12-27"
	dp.AlumniDate = "2021-01-01"
	return dp, nil
}

func (dp MockDiplomaImpl) CheckDiploma() bool {
	return true
}

func (dp MockDiplomaImpl) LogFields() log.Fields {
	return log.Fields{"first_name": dp.FirstName, "last_name": dp.LastName, "birth_date": dp.BirthDate, "alumni_date": dp.AlumniDate}
}

func (dp MockDiplomaImpl) String() string {
	str := dp.FirstName + ", " + dp.LastName + ", " + dp.BirthDate + ", " + dp.AlumniDate
	return str
}

func (dp MockDiplomaImpl) AddToRetry() {
	//Impl me
}

func (dp MockDiplomaImpl) EthWriting() (string, bool) {
	return "", true
}

func (dp MockDiplomaImpl) EthGetter() (float64, [30]api.Skill, error) {
 	return 0, [30]api.Skill{}, errors.New("")
}
