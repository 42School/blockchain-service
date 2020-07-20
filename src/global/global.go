package global

import (
	"container/list"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

var PasswordAccount string = os.Getenv("KEYPASSWD")
var PathKeyStore string = os.Getenv("KEYSTOREPATH")
var FileKeyStore string = os.Getenv("KEYSTOREFILE")
var OfficialAddress common.Address = common.HexToAddress(os.Getenv("OFFICIALADDRESS"))
var NetworkLink string = os.Getenv("NETWORKLINK")
var AddressOfContract string = os.Getenv("ADDRESSCONTRACT")

var FtEndPoint string = os.Getenv("FTENDPOINT")

var Env string = os.Getenv("RUNENV")

var DevAddress string = os.Getenv("DEVADDRESS")
var DevPrivateKey string = os.Getenv("DEVPRIVATEKEY")

var ToCheckHash *list.List = list.New()