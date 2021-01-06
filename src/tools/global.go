package tools

import (
	"container/list"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"context"
)

var PasswordAccount string = os.Getenv("KEYPASSWD")
var PathKeyStore string = os.Getenv("KEYSTOREPATH")
var PathKeyStoreSign string = os.Getenv("KEYSTOREPATHSIGN")
var NetworkLink string = os.Getenv("NETWORKLINK")
var AddressOfContract string = os.Getenv("ADDRESSCONTRACT")

var FtEndPoint string = os.Getenv("FTENDPOINT")
var ValidationPath string = os.Getenv("VALIDATIONPATH")
var RetryPath string = os.Getenv("RETRYPATH")
var Token string = os.Getenv("TOKEN")

var Env string = os.Getenv("RUNENV")

var SecuritySystem bool = false

var RetryQueue *list.List = list.New()
var ToCheckHash *list.List = list.New()

var RetryDB		*mongo.Collection
var ToCheckDB	*mongo.Collection
var CtxDB		context.Context