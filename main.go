package main

import (
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/rest"
	"github.com/42School/blockchain-service/src/async"
	"github.com/42School/blockchain-service/src/tools"
	"log"
	"net/http"
)

func main() {
	go async.ValideHash()
	go async.RetryDiploma()
	go async.ReadStdin()
	tools.LogsMsg("Blockchain Service is running !")
	account.CreateAccountsManager()
	router := rest.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
