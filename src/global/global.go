package global

import (
	"container/list"
	"os"
)

var PasswordAccount string = os.Getenv("KEYPASSWD")
var PathKeyStore string = os.Getenv("KEYSTOREPATH")
var PathKeyStoreSign string = os.Getenv("KEYSTOREPATHSIGN")
var NetworkLink string = os.Getenv("NETWORKLINK")
var AddressOfContract string = os.Getenv("ADDRESSCONTRACT")

var FtEndPoint string = os.Getenv("FTENDPOINT")
var PathValidation string = os.Getenv("VALIDATIONPATH")
var PathRetry string = os.Getenv("PATHRETRY")

var Env string = os.Getenv("RUNENV")

var SecuritySystem bool = false

var RetryQueue *list.List = list.New()
var ToCheckHash *list.List = list.New()