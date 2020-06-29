package controllers

import (
	"encoding/json"
	"github.com/lpieri/42-Diploma/src/api/models"
	"io/ioutil"
	"log"
	"net/http"
)


func CreateDiploma(w http.ResponseWriter, r *http.Request) {
	var newDiploma models.Diploma
	jsonData, readErr := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(jsonData, &newDiploma)
	if r.ContentLength == 0 || readErr != nil || jsonErr != nil || models.CheckDiploma(newDiploma) == false {
		log.Println("Request Fail !!")
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("In CreateDiploma but fail")
		return
	}
	if models.NewDiploma(newDiploma) != true {
		log.Println("Request Fail !!")
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("In CreateDiploma but fail")
		return
	}
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("In CreateDiploma but success")
	return
}

func CheckHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Print(r.Body)
	json.NewEncoder(w).Encode("In CheckHash")
}