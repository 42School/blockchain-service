package diplomas

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type DiplomaImplMock struct {
	Id         uuid.UUID `bson:"_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	AlumniDate string    `json:"alumni_date"`
	Level      float64   `json:"level"`
	Skills     []Skill   `json:"skills"`
	Counter    int       `json:"counter"`
}

func (dp DiplomaImplMock) CheckDiploma() bool {
	return true
}

func (dp DiplomaImplMock) LogFields() log.Fields {
	return log.Fields{"first_name": dp.FirstName, "last_name": dp.LastName, "birth_date": dp.BirthDate, "alumni_date": dp.AlumniDate}
}

func (dp DiplomaImplMock) String() string {
	str := dp.FirstName + ", " + dp.LastName + ", " + dp.BirthDate + ", " + dp.AlumniDate
	return str
}

func (dp DiplomaImplMock) AddToRetry() {
	//Impl me
}

func (dp DiplomaImplMock) EthWriting() (string, bool) {
	return "", true
}

func (dp DiplomaImplMock) EthGetter() (float64, [30]Skill, error) {
 	return 0, [30]Skill{}, errors.New("")
}
