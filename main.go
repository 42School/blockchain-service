package main

import (
	"log"
	"net/http"

	"github.com/lpieri/42-Diploma/src/api"
)

func main() {
	router := api.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
