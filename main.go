package main

import (
	"github.com/lpieri/42-Diploma/src/account"
	"log"
	"net/http"

	"github.com/lpieri/42-Diploma/src/api"
)

func main() {
	account.CreateAccountsManager()
	router := api.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
