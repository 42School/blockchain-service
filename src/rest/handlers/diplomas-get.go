package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"net/http"
)

type GetDiplomaHandler struct {
	diploma diplomas.Diploma
	err     error
}

func NewGetDiplomaHandler() *GetDiplomaHandler {
	var u = GetDiplomaHandler{diplomas.NewDiploma(), errors.New("")}
	return &u
}

func (uHandler *GetDiplomaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	uHandler.diploma, uHandler.err = uHandler.diploma.ReadJson(r.Body)
	if uHandler.err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	level, skills, err := uHandler.diploma.EthGetter()
	if err != nil {
		http.Error(w, "The request is fail, please retry & check the data.", http.StatusBadRequest)
		return
	} else {
		res, _ := json.Marshal(ResponseJson{true, "", ResponseData{"", level, skills[:]}})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}
