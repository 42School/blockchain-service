package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/dao/interfaces"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CreateDiplomaHandler struct {
	diploma  interfaces.Diploma
	err      error
}

func NewCreateDiplomaHandler() *CreateDiplomaHandler {
	var u = CreateDiplomaHandler{interfaces.NewDiploma(), errors.New("")}
	return &u
}

func (uHandler *CreateDiplomaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if r.Header.Get("Token") != tools.Token {
		http.Error(w, "You are not authorized !", http.StatusUnauthorized)
		return
	}
	uHandler.diploma, uHandler.err = api.WebhookToDiploma(r.Body)
	if uHandler.err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	if uHandler.diploma.CheckDiploma() == false {
		http.Error(w, "The data sent are not valid, to be written in blockchain please try again !", http.StatusBadRequest)
		return
	}
	log.WithFields(uHandler.diploma.LogFields()).Debug("Received new request to write diploma.")
	hash, bool := uHandler.diploma.EthWriting()
	if bool == false {
		http.Error(w, "Blockchain writing had a problem, the diploma is saved in the queue.", http.StatusBadRequest)
		return
	} else {
		res, _ := json.Marshal(ResponseJson{true, "The writing in blockchain has been done, it will be confirmed in 10 min.", ResponseData{hash, 0, []diplomas.Skill{}}})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}
