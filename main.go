package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lpieri/42-Diploma/src/account"
	"github.com/lpieri/42-Diploma/src/api"
	"github.com/lpieri/42-Diploma/src/contracts"
	"github.com/lpieri/42-Diploma/src/global"
	"log"
	"net/http"
	"time"
)

func ValidedHash() {
	for {
		time.Sleep(600000 * time.Millisecond)
		copyList := global.ToCheckHash
		for e := copyList.Front(); e != nil; e = copyList.Front() {
			if e != nil {
				hash, _ := e.Value.([]byte)
				_, _, err := contracts.CallGetDiploma(hash)
				if err == nil {
					// Send http request
					strHash := hexutil.Encode(hash)
					log.Println("Valide write.", strHash)
					global.ToCheckHash.Remove(e)
				}
			}
		}
	}
}

func main() {
	go ValidedHash()
	account.CreateAccountsManager()
	router := api.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
