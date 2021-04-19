package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/dao/interfaces"
	"net/http"
)

type GetDiplomaHandler struct {
	diploma  interfaces.Diploma
	err      error
}

func NewGetDiplomaHandler() *GetDiplomaHandler {
	var u = GetDiplomaHandler{interfaces.NewDiploma(), errors.New("")}
	return &u
}

func (uHandler *GetDiplomaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var diplomaData diplomas.DiplomaImpl
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&diplomaData)
	uHandler.diploma = diplomaData
	if err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	level, skills, errGet := uHandler.diploma.EthGetter()
	if errGet != nil {
		http.Error(w, "The request is fail, please retry & check the data.", http.StatusBadRequest)
		return
	} else {
		res, _ := json.Marshal(ResponseJson{true, "", ResponseData{"", level, skills[:]}})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}
