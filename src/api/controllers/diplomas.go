package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func test() {
	log.Println("test import controllers")
}

func Print(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Coucou toi je suis dans controllers !")
}
