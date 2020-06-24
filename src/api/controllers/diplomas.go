package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)


func CreateDiploma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Print(r.Body)
	json.NewEncoder(w).Encode("In CreateDiploma")
}

func CheckHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Print(r.Body)
	json.NewEncoder(w).Encode("In CheckHash")
}