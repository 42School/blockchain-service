package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/tools"
	"net/http"
)

type GetAllDiplomaHandler struct {
	err error
}

func NewGetAllDiplomaHandler() *GetAllDiplomaHandler {
	var u = GetAllDiplomaHandler{ errors.New("")}
	return &u
}

func (uHandler *GetAllDiplomaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if r.Header.Get("Token") != tools.Token {
		http.Error(w, "You are not authorized !", http.StatusUnauthorized)
		return
	}
	diplomas := diplomas.EthAllGetter()
	if diplomas == nil {
		http.Error(w, "A problem occurred during data recovery.", http.StatusInternalServerError)
		return
	}
	uHandler.err = json.NewEncoder(w).Encode(diplomas)
	if uHandler.err != nil {
		http.Error(w, "Fail Encode json.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
