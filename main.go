package main

import (
	"github.com/42School/blockchain-service/src/api/models"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/api"
	"github.com/42School/blockchain-service/src/contracts"
	"github.com/42School/blockchain-service/src/global"
	"log"
	"net/http"
	"strings"
	"time"
)

func ValidedHash() {
	for {
		time.Sleep(10 * time.Minute)
		copyList := global.ToCheckHash
		for e := copyList.Front(); e != nil; {
			if e != nil {
				hash, _ := e.Value.([]byte)
				_, _, err := contracts.CallGetDiploma(hash)
				if err == nil {
					strHash := hexutil.Encode(hash)
					data := "{'Status': true, 'Message': 'The " + strHash + " diploma is definitely inscribed on Ethereum.', 'Data': {" + strHash + "}}"
					_, err := http.Post(global.FtEndPoint + "/check-request", "Content-Type: application/json", strings.NewReader(data))
					if err == nil {
						global.ToCheckHash.Remove(e)
						e = copyList.Front()
					} else {
						e = e.Next()
					}
				} else {
					e = e.Next()
				}
			}
		}
	}
}

func RetryDiploma () {
	for {
		time.Sleep(30 * time.Minute)
		copyList := global.RetryQueue
		for e := copyList.Front(); e != nil; {
			if e != nil {
				diploma, _ := e.Value.(models.Diploma)
				tools.LogsDev(diploma.String())
				hash, bool := diploma.EthWriting()
				if bool == true {
					data := "{'Status':true,'Message':'The writing in blockchain has been done, it will be confirmed in 10 min.','Data':{'Hash': " + hash + ",'Level':0,'Skills':[]}}"
					http.Post(global.FtEndPoint + "/check-request", "Content-Type: application/json", strings.NewReader(data))
					e = copyList.Front()
					global.RetryQueue.Remove(e)
				} else {
					e = e.Next()
				}
			}
		}
	}
}

func main() {
	go ValidedHash()
	go RetryDiploma()
	tools.LogsMsg("Blockchain Service is running !")
	account.CreateAccountsManager()
	router := api.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
