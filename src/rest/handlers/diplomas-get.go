package handlers

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"net/http"
)

func GetDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	var diploma diplomas.Diploma
	err := json.NewDecoder(r.Body).Decode(&diploma)
	if err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	level, skills, errGet := diploma.EthGetter()
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
