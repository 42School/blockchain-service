package controllers

import (
	"encoding/json"
	"github.com/lpieri/42-Diploma/src/api/models"
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
	jsonData, readErr := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(jsonData, &newDiploma)
	if r.ContentLength == 0 || readErr != nil || jsonErr != nil || models.CheckDiploma(newDiploma) == false {
		res, _ := json.Marshal(ResponseJson{false, "The data sent are not valid, to be written in blockchain please try again!", ResponseData{}})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	var res []byte
	hash, bool := models.NewDiploma(newDiploma)
	if bool == false {
		res, _ = json.Marshal(ResponseJson{true, "Blockchain writing had a problem, please try again.", ResponseData{"", 0, []float64{}}})
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ = json.Marshal(ResponseJson{true, "The writing in blockchain has been done, it will be confirmed in 10 min.", ResponseData{hash, 0, []float64{}}})
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
	return
}

func GetDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	var diploma models.Diploma
	var res []byte
	jsonData, readErr := ioutil.ReadAll(r.Body)
	jsonErr := json.Unmarshal(jsonData, &diploma)
	level, skills, errGet := models.GetDiploma(diploma)
	if r.ContentLength == 0 || readErr != nil || jsonErr != nil || errGet != nil {
		log.Println("Request Fail !!")
		res, _ = json.Marshal(ResponseJson{false, "The request is fail, please retry & check the data", ResponseData{}})
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ = json.Marshal(ResponseJson{true, "", ResponseData{"", level, skills[:]}})
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
	return
}
