package global

import (
	//"github.com/lpieri/42-Diploma/src/queue"
	"container/list"
	"os"
)

var PasswordAccount string = os.Getenv("KEYPASSWD")
var PathKeyStore string = os.Getenv("KEYSTORE")
var NetworkLink string = os.Getenv("NETWORKLINK")
var AddressOfContract string = os.Getenv("ADDRESSCONTRACT")

//var RetryQueue *queue.Queue = queue.New()
var ToCheckHash *list.List = list.New()