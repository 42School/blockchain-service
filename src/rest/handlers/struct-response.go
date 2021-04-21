package handlers

import (
	"github.com/42School/blockchain-service/src/dao/api"
)

type ResponseData struct {
	Hash   string
	Level  float64
	Skills []api.Skill
}

type ResponseJson struct {
	Status  bool
	Message string
	Data    ResponseData
}
