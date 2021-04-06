package handlers

import "github.com/42School/blockchain-service/src/dao/diplomas"

type ResponseData struct {
	Hash	string
	Level	float64
	Skills	[]diplomas.Skill
}

type ResponseJson struct {
	Status	bool
	Message	string
	Data	ResponseData
}
