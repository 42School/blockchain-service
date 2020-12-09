package handlers

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/global"
	"github.com/42School/blockchain-service/src/tools"
	"net/http"
)

func CreateDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	var newDiploma diplomas.Diploma
	if r.Header.Get("Token") != global.Token {
		http.Error(w, "You are not authorized !", http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newDiploma)
	if err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	if newDiploma.CheckDiploma() == false {
		http.Error(w, "The data sent are not valid, to be written in blockchain please try again !", http.StatusBadRequest)
		return
	}
	if global.SecuritySystem {
		newDiploma.AddToRetry()
		http.Error(w, "The security system is activated, the request has just been queued.", http.StatusInternalServerError)
		return
	}
	tools.LogsDev("Received request to write the " + newDiploma.FirstName + " " + newDiploma.LastName + " diploma.")
	hash, bool := newDiploma.EthWriting()
	if bool == false {
		http.Error(w, "Blockchain writing had a problem, the diploma is saved in the queue.", http.StatusBadRequest)
		return
	} else {
		res, _ := json.Marshal(ResponseJson{true, "The writing in blockchain has been done, it will be confirmed in 10 min.", ResponseData{hash, 0, []float64{}}})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}