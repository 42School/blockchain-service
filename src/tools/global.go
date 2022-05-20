package tools

import (
	"container/list"
	"github.com/42School/blockchain-service/src/db"
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

// Other Variable
var Env = os.Getenv("RUN_ENV")

var RetryQueue = list.New()
var ToCheckHash = list.New()

var Db db.Database
