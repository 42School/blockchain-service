package interfaces

import (
	"github.com/42School/blockchain-service/src/dao/diplomas"
	log "github.com/sirupsen/logrus"
)

type Diploma interface {
	CheckDiploma() bool
	LogFields() log.Fields
	AddToRetry()
	String() string
	EthWriting() (string, bool)
	EthGetter() (float64, [30]diplomas.Skill, error)
}

func NewDiploma() Diploma {
	var i Diploma
	i = diplomas.DiplomaImpl{}
	return i
}