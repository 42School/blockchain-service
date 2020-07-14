package global

import (
	"container/list"
	"os"
)

var PasswordAccount string = os.Getenv("KEYPASSWD")
var PathKeyStore string = os.Getenv("KEYSTORE")
var NetworkLink string = os.Getenv("NETWORKLINK")
var AddressOfContract string = os.Getenv("ADDRESSCONTRACT")

var Env string = os.Getenv("RUNENV")

var DevAddress string = os.Getenv("DEVADDRESS")
var DevPrivateKey string = os.Getenv("DEVPRIVATEKEY")

var ToCheckHash *list.List = list.New()