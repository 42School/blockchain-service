package interfaces

import (
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	log "github.com/sirupsen/logrus"
	"io"
)

type Diploma interface {
	CheckDiploma() bool
	LogFields() log.Fields
	AddToRetry()
	String() string
	EthWriting() (string, bool)
	EthGetter() (float64, [30]api.Skill, error)
	ReadWebhook(body io.ReadCloser) (diplomas.DiplomaImpl, error)
	ReadJson(body io.ReadCloser) (diplomas.DiplomaImpl, error)
}

func NewDiploma() Diploma {
	var i Diploma
	i = &diplomas.DiplomaImpl{}
	return i
}