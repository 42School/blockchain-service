package mocks

import (
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"io"
)

type MockDiplomaImpl struct {
	mock.Mock
}

func (dp MockDiplomaImpl) ReadWebhook(body io.ReadCloser) (diplomas.Diploma, error) {
	args := dp.Called()
	return args.Get(0).(diplomas.Diploma), args.Error(1)
}

func (dp MockDiplomaImpl) ReadJson(body io.ReadCloser) (diplomas.Diploma, error) {
	args := dp.Called()
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
