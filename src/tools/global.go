package tools

import (
	"container/list"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// Keystore File
var PathKeyStoreSign string = os.Getenv("KEYSTORE_PATH_SIGN")
var PasswordAccount string = os.Getenv("KEY_PASSWD")
var PathKeyStore string = os.Getenv("KEYSTORE_PATH")
var AccountsFile string = os.Getenv("ACCOUNTS_FILE")

// Blockchain Variable
var NetworkLink string = os.Getenv("NETWORK_LINK")
var AddressOfContract string = os.Getenv("CONTRACT_ADDRESS")

// Endpoint of 42 intra
var FtEndPoint string = os.Getenv("FT_END_POINT")
var RetryPath string = os.Getenv("RETRY_PATH")
var ValidationPath string = os.Getenv("VALIDATION_PATH")
var Token string = os.Getenv("TOKEN")

// Mongo Variable
var MongoIp string = os.Getenv("MONGO_IP")
var MongoPort string = os.Getenv("MONGO_PORT")
var MongoUser string = os.Getenv("MONGO_USER")
var MongoPasswd string = os.Getenv("MONGO_PASSWD")

// Other Variable
var Env string = os.Getenv("RUN_ENV")

var SecuritySystem bool = false

var RetryQueue *list.List = list.New()
var ToCheckHash *list.List = list.New()

var RetryDB		*mongo.Collection
var ToCheckDB	*mongo.Collection