package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/tools"
	"net/http"
)

type GetAllDiplomaHandler struct {
	blockchain contracts.BlockchainFunc
	err error
}

func NewGetAllDiplomaHandler() *GetAllDiplomaHandler {
	var u = GetAllDiplomaHandler{contracts.NewBlockchainFunc(), errors.New("")}
	return &u
}

func (uHandler *GetAllDiplomaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if r.Header.Get("Token") != tools.Token {
		http.Error(w, "You are not authorized !", http.StatusUnauthorized)
		return
	}
	diplomas, err := uHandler.blockchain.CallGetAllDiploma()
	if err != nil {
		tools.LogsError(err)
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
