package diplomas

import (
	"github.com/42School/blockchain-service/src/dao/api"
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
	ReadWebhook(body io.ReadCloser) (Diploma, error)
	ReadJson(body io.ReadCloser) (Diploma, error)
}

func NewDiploma() Diploma {
	var i Diploma
	i = &DiplomaImpl{}
	return i
}