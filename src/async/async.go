package async

import (
	"bufio"
	"fmt"
	"github.com/42School/blockchain-service/src/api/models"
	"github.com/42School/blockchain-service/src/contracts"
	"github.com/42School/blockchain-service/src/global"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func ValideHash() {
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
					_, err := http.Post(global.FtEndPoint + global.ValidationPath, "Content-Type: application/json", strings.NewReader(data))
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
		if global.SecuritySystem == false {
			copyList := global.RetryQueue
			for e := copyList.Front(); e != nil; {
				if e != nil {
					diploma, _ := e.Value.(models.Diploma)
					tools.LogsDev(diploma.String())
					hash, bool := diploma.EthWriting()
					if bool == true {
						data := "{'Status':true,'Message':'The writing in blockchain has been done, it will be confirmed in 10 min.','Data':{'Hash': " + hash + ",'Level':0,'Skills':[]}}"
						http.Post(global.FtEndPoint + global.RetryPath, "Content-Type: application/json", strings.NewReader(data))
						global.RetryQueue.Remove(e)
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
				global.SecuritySystem = false
				fmt.Println("The security system has been disabled !")
				fmt.Println("Goodbye of cmd mode")
			} else if text == "Exit" || text == "exit" || text == "EXIT" {
				cmd = false
				fmt.Println("Goodbye of cmd mode")
			}
		}
	}
}