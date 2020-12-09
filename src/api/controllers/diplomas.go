package controllers

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/api/models"
	"github.com/42School/blockchain-service/src/global"
	"github.com/42School/blockchain-service/src/tools"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseData struct {
	Hash	string
	Level	float64
	Skills	[]float64
}

type ResponseJson struct {
	Status	bool
	Message	string
	Data	ResponseData
}

func CreateDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	var newDiploma models.Diploma
	if r.Header.Get("Token") != global.Token {
		http.Error(w, "You are not authorized !", http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newDiploma)
	if err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	if r.ContentLength == 0 || newDiploma.CheckDiploma() == false {
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

func GetDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	var diploma models.Diploma
	err := json.NewDecoder(r.Body).Decode(&diploma)
	if err != nil {
		http.Error(w, "Fail Unmarshalling json", http.StatusBadRequest)
		return
	}
	level, skills, errGet := diploma.EthGetter()
	if r.ContentLength == 0 || errGet != nil {
		http.Error(w, "The request is fail, please retry & check the data.", http.StatusBadRequest)
		return
	} else {
		res, _ := json.Marshal(ResponseJson{true, "", ResponseData{"", level, skills[:]}})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	return
}

func CheckRouter (w http.ResponseWriter, r *http.Request) {
	jsonData, _ := ioutil.ReadAll(r.Body)
	log.Println("r.body", string(jsonData))
}