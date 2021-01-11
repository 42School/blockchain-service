package handlers

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/tools"
	"net/http"
)

func GetAllDiplomas(w http.ResponseWriter, r *http.Request) {
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
	err := json.NewEncoder(w).Encode(diplomas)
	if err != nil {
		http.Error(w, "Fail Encode json.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
