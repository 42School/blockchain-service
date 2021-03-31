package async

import (
	"context"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func ValideHash() {
	url := tools.FtEndPoint + tools.ValidationPath
	for {
		time.Sleep(10 * time.Minute)
		copyList := tools.ToCheckHash
		for e := copyList.Front(); e != nil; {
			if e != nil {
				check, _ := e.Value.(diplomas.VerificationHash)
				strHash := hexutil.Encode(check.StudentHash)
				data := "{'Status': true, 'Message': 'The " + strHash + " diploma is definitely inscribed on Ethereum.', 'Data': {" + strHash + "}}"
				client, _ := ethclient.Dial(tools.NetworkLink)
				receipt, err := client.TransactionReceipt(context.Background(), check.Tx.Hash())
				if err == nil {
					if receipt.Status == 1 {
						contracts.CheckSecurity(client, check.Tx, check.StudentHash)
						_, err = http.Post(url, "Content-Type: application/json", strings.NewReader(data))
						if err == nil {
							tools.ToCheckHash.Remove(e)
							txByte, _ := check.Tx.MarshalJSON()
							tools.ToCheckDB.DeleteOne(context.TODO(), bson.M{"tx": txByte, "studenthash": check.StudentHash})
							e = copyList.Front()
							continue
						}
					} else {
						revertMsg := contracts.GetRevert(client, check.Tx, receipt)
						if revertMsg != "" {
							if strings.Contains(revertMsg, "FtDiploma: Is not 42 sign this diploma") {
								data = "{'Status': false, 'Message': 'The " + strHash + " diploma wasn't signed by 42, so it's not in the blockchain.', 'Data': {" + strHash + "}}"
							}
							_, err = http.Post(url, "Content-Type: application/json", strings.NewReader(data))
							if err == nil {
								tools.ToCheckHash.Remove(e)
								txByte, _ := check.Tx.MarshalJSON()
								tools.ToCheckDB.DeleteOne(context.TODO(), bson.M{"tx": txByte, "studenthash": check.StudentHash})
								e = copyList.Front()
								continue
							}
						}
					}
				}
				e = e.Next()
			}
		}
	}
}

func RetryDiploma () {
	url := tools.FtEndPoint + tools.RetryPath
	for {
		time.Sleep(30 * time.Minute)
		copyList := tools.RetryQueue
		for e := copyList.Front(); e != nil; {
			if e != nil {
				diploma, _ := e.Value.(diplomas.Diploma)
				log.WithFields(diploma.LogFields()).Debug("Try to retry a diploma")
				hash, bool := diploma.EthWriting()
				if bool == true {
					data := "{'Status':true,'Message':'The writing in blockchain has been done, it will be confirmed in 10 min.','Data':{'Hash': " + hash + ",'Level':0,'Skills':[]}}"
					http.Post(url, "Content-Type: application/json", strings.NewReader(data))
					tools.RetryQueue.Remove(e)
					tools.RetryDB.DeleteOne(context.TODO(), diploma)
					e = copyList.Front()
				} else {
					e = e.Next()
				}
			}
		}
	}
}
