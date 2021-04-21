package tools

import (
	"container/list"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// Keystore File
var PathKeyStoreSign = os.Getenv("KEYSTORE_PATH_SIGN")
var PasswordAccount = os.Getenv("KEY_PASSWD")
var PathKeyStore = os.Getenv("KEYSTORE_PATH")
var AccountsFile = os.Getenv("ACCOUNTS_FILE")

// Blockchain Variable
var NetworkLink = os.Getenv("NETWORK_LINK")
var AddressOfContract = os.Getenv("CONTRACT_ADDRESS")

// Endpoint of 42 intra
var FtEndPoint = os.Getenv("FT_END_POINT")
var RetryPath = os.Getenv("RETRY_PATH")
var ValidationPath = os.Getenv("VALIDATION_PATH")
var Token = os.Getenv("TOKEN")

// Mongo Variable
var MongoIp = os.Getenv("MONGO_IP")
var MongoPort = os.Getenv("MONGO_PORT")
var MongoUser = os.Getenv("MONGO_USER")
var MongoPasswd = os.Getenv("MONGO_PASSWD")

// Other Variable
var Env = os.Getenv("RUN_ENV")

var RetryQueue = list.New()
var ToCheckHash = list.New()

var RetryDB *mongo.Collection
var ToCheckDB *mongo.Collection
