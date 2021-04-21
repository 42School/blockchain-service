package mocks

import (
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"io"
)

type MockDiplomaImpl struct {
	mock.Mock
	Id         uuid.UUID `bson:"_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	AlumniDate string    `json:"alumni_date"`
	Level      float64   `json:"level"`
	Skills     []api.Skill   `json:"skills"`
	Counter    int       `json:"counter"`
}

func (dp MockDiplomaImpl) ReadWebhook(body io.ReadCloser) (diplomas.Diploma, error) {
	args := dp.Called()
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
	return args.Get(0).(diplomas.Diploma), args.Error(1)
}

func (dp MockDiplomaImpl) ReadJson(body io.ReadCloser) (diplomas.Diploma, error) {
	args := dp.Called()
	dp.FirstName = "Louise"
	dp.LastName = "Pieri"
	dp.BirthDate = "1998-12-27"
	dp.AlumniDate = "2021-01-01"
	return args.Get(0).(diplomas.Diploma), args.Error(1)
}

func (dp MockDiplomaImpl) CheckDiploma() bool {
	args := dp.Called()
	return args.Bool(0)
}

func (dp MockDiplomaImpl) LogFields() log.Fields {
	return log.Fields{"first_name": "Louise", "last_name": "Pieri", "birth_date": "1998-12-27", "alumni_date": "2021-01-01"}
}

func (dp MockDiplomaImpl) String() string {
	args := dp.Called()
	return args.String(0)
}

func (dp MockDiplomaImpl) AddToRetry() {
	//Impl me
}

func (dp MockDiplomaImpl) EthWriting() (string, bool) {
	args := dp.Called()
	return args.String(0), args.Bool(1)
}

func (dp MockDiplomaImpl) EthGetter() (float64, [30]api.Skill, error) {
	args := dp.Called()
 	return float64(args.Int(0)), args.Get(1).([30]api.Skill), args.Error(2)
}
