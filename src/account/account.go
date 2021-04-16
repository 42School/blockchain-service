package account

import (
	"crypto/ecdsa"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type Account struct {
	KeyStoreFile string
	Password     string
}

var KeyStore *keystore.KeyStore
var CurrentAccount = 0
var Accounts []Account

func ParseAccounts() {
	bits, err := ioutil.ReadFile(tools.AccountsFile)
	if err == nil {
		data := string(bits)
		lines := strings.Split(data, "\n")
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if line != "" {
				if line[0] != '#' {
					accountData := strings.Split(line, ", ")
					account := Account{accountData[0], accountData[1]}
					Accounts = append(Accounts, account)
				}
			}
		}
	}
}

func CreateAccountsManager() {
	KeyStore = keystore.NewKeyStore(tools.PathKeyStoreSign, keystore.StandardScryptN, keystore.StandardScryptP)
	ParseAccounts()
}

func GetSignAccount() accounts.Account {
	return KeyStore.Accounts()[0]
}

func GetWriterAccount() (common.Address, *ecdsa.PrivateKey, error) {
	keyjson, errRead := ioutil.ReadFile(tools.PathKeyStore + "/" + Accounts[CurrentAccount].KeyStoreFile)
	if errRead != nil {
		return common.Address{}, nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, Accounts[CurrentAccount].Password)
	if errDecrypt != nil {
		return common.Address{}, nil, errDecrypt
	}
	log.WithFields(log.Fields{"private_key": hexutil.Encode(crypto.FromECDSA(key.PrivateKey))}).Debug("The private key of the wallet writer")
	return key.Address, key.PrivateKey, nil
}

func ChangeAccount() {
	// Send Mail of Current Account
	address, _, _ := GetWriterAccount()
	tools.SendMail("Empty Account", address.Hex())
	if CurrentAccount+1 == len(Accounts) {
		CurrentAccount = 0
	} else {
		CurrentAccount = CurrentAccount + 1
	}
}
