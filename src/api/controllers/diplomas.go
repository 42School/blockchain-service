package controllers

import (
	"encoding/json"
	"github.com/lpieri/42-Diploma/src/api/models"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseData struct {
	Level	float64
	Skills	[]float64
}

type ResponseJson struct {
	Status	bool
	Message	string
	Data	ResponseData
}

func CreateDiploma(w http.ResponseWriter, r *http.Request) {
	var newDiploma models.Diploma
	jsonData, readErr := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(jsonData, &newDiploma)
	if r.ContentLength == 0 || readErr != nil || jsonErr != nil || models.CheckDiploma(newDiploma) == false {
		//log.Println("Request Fail !!")
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("In CreateDiploma but fail")
		return
	}
	if models.NewDiploma(newDiploma) != true {
		//log.Println("Request Fail !!")
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

func GetDiploma(w http.ResponseWriter, r *http.Request) {
	var diploma models.Diploma
	jsonData, readErr := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(jsonData, &diploma)
	level, skills, errGet := models.GetDiploma(diploma)
	if r.ContentLength == 0 || readErr != nil || jsonErr != nil || errGet != nil {
		log.Println("Request Fail !!")
		response := ResponseJson{false, "The request is fail, please retry & check the data", nil}
		res, _ := json.Marshal(response)
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	response := ResponseJson{true, "", ResponseData{level, skills[:]}}
	res, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func CheckHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Print(r.Body)
	json.NewEncoder(w).Encode("In CheckHash")
}