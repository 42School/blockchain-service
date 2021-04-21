package mocks

import (
	"github.com/42School/blockchain-service/src/dao/diplomas"
)

func NewMockDiploma() diplomas.Diploma {
	var i diplomas.Diploma
	i = MockDiplomaImpl{}
	return i
}