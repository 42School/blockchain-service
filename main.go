package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lpieri/42-Diploma/src/account"
	"github.com/lpieri/42-Diploma/src/api"
	"github.com/lpieri/42-Diploma/src/contracts"
	"github.com/lpieri/42-Diploma/src/global"
	"log"
	"net/http"
	"strings"
	"time"
)

func ValidedHash() {
	for {
		time.Sleep(6000/*00*/ * time.Millisecond)
		copyList := global.ToCheckHash
		for e := copyList.Front(); e != nil; e = copyList.Front() {
			if e != nil {
				hash, _ := e.Value.([]byte)
				_, _, err := contracts.CallGetDiploma(hash)
				if err == nil {
					strHash := hexutil.Encode(hash)
					data := "{'Status': true, 'Message': 'The " + strHash + " diploma is definitely inscribed on Ethereum.', 'Data': {" + strHash + "}}"
					_, err := http.Post("http://end-point.42.fr/path", "Content-Type: application/json", strings.NewReader(data))
					if err == nil {
						global.ToCheckHash.Remove(e)
					}
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
