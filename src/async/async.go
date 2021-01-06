package async

import (
	"bufio"
	"context"
	"fmt"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func ValideHash() {
	url := tools.FtEndPoint + tools.ValidationPath
	for {
		time.Sleep(1 * time.Minute)
		copyList := tools.ToCheckHash
		for e := copyList.Front(); e != nil; {
			if e != nil {
				check, _ := e.Value.(diplomas.VerificationHash)
				strHash := hexutil.Encode(check.StudentHash)
				client, _ := ethclient.Dial(tools.NetworkLink)
				receipt, err := client.TransactionReceipt(context.Background(), check.Tx.Hash())
				if err == nil {
					if receipt.Status == 1 {
						contracts.CheckSecurity(client, check.Tx, check.StudentHash)
						data := "{'Status': true, 'Message': 'The " + strHash + " diploma is definitely inscribed on Ethereum.', 'Data': {" + strHash + "}}"
						_, err := http.Post(url, "Content-Type: application/json", strings.NewReader(data))
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
							data := ""
							if strings.Contains(revertMsg, "FtDiploma: Is not 42 sign this diploma") {
								data = "{'Status': false, 'Message': 'The " + strHash + " diploma wasn't signed by 42, so it's not in the blockchain.', 'Data': {" + strHash + "}}"
							} else if strings.Contains(revertMsg, "FtDiploma: The diploma already exists.") {
								data = "{'Status': true, 'Message': 'The " + strHash + " diploma is definitely inscribed on Ethereum.', 'Data': {" + strHash + "}}"
							}
							_, err := http.Post(url, "Content-Type: application/json", strings.NewReader(data))
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
		time.Sleep(1 * time.Minute)
		if tools.SecuritySystem == false {
			copyList := tools.RetryQueue
			for e := copyList.Front(); e != nil; {
				if e != nil {
					diploma, _ := e.Value.(diplomas.Diploma)
					tools.LogsDev(diploma.String())
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
}

func ReadStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	var cmd = false
	for {
		if cmd {
			fmt.Print("> ")
		}
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Println(err)
		} else {
			text := scanner.Text()
			if text == "CMD" || text == "cmd" || text == "Cmd"{
				fmt.Println("Please enter your command:")
				fmt.Println(" - 'disable security system': to disable the security system")
				fmt.Println(" - 'exit' or 'Exit' or 'EXIT': to exit the CMD mode")
				cmd = true
			} else if text == "disable security system" && cmd {
				cmd = false
				tools.SecuritySystem = false
				fmt.Println("The security system has been disabled !")
				fmt.Println("Goodbye of cmd mode")
			} else if text == "Exit" || text == "exit" || text == "EXIT" {
				cmd = false
				fmt.Println("Goodbye of cmd mode")
			}
		}
	}
}