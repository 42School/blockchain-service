package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CheckRouter (w http.ResponseWriter, r *http.Request) {
	jsonData, _ := ioutil.ReadAll(r.Body)
	log.Println("r.body", string(jsonData))
}
