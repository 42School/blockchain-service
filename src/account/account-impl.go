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

type AccountsImpl struct {
	KeyStore *keystore.KeyStore
	CurrentAccount int
	Accounts []Account
}

var Accounts AccountsManager

func (ac *AccountsImpl) InitAccounts() {
	ac.KeyStore = keystore.NewKeyStore(tools.PathKeyStoreSign, keystore.StandardScryptN, keystore.StandardScryptP)
	ac.CurrentAccount = 0
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
					ac.Accounts = append(ac.Accounts, account)
				}
			}
		}
	}
}

func (ac *AccountsImpl) GetLenAccounts() int {
	return len(ac.Accounts)
}

func (ac *AccountsImpl) ChangeWriter() {
	if ac.CurrentAccount+1 == len(ac.Accounts) {
		ac.CurrentAccount = 0
	} else {
		ac.CurrentAccount = ac.CurrentAccount + 1
	}
}

func (ac *AccountsImpl) GetSign() accounts.Account {
	return ac.KeyStore.Accounts()[0]
}

func (ac *AccountsImpl) SignHash(hash common.Hash) ([]byte, error) {
	sign, err := ac.KeyStore.SignHashWithPassphrase(ac.GetSign(), tools.PasswordAccount, hash.Bytes())
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func (ac *AccountsImpl) GetWriter() (common.Address, *ecdsa.PrivateKey, error) {
	keyjson, errRead := ioutil.ReadFile(tools.PathKeyStore + "/" + ac.Accounts[ac.CurrentAccount].KeyStoreFile)
	if errRead != nil {
		return common.Address{}, nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, ac.Accounts[ac.CurrentAccount].Password)
	if errDecrypt != nil {
		return common.Address{}, nil, errDecrypt
	}
	log.WithFields(log.Fields{"private_key": hexutil.Encode(crypto.FromECDSA(key.PrivateKey))}).Debug("The private key of the wallet writer")
	return key.Address, key.PrivateKey, nil
}

func (ac *AccountsImpl) GetWriterByI(i int) (common.Address, *ecdsa.PrivateKey, error) {
	keyjson, errRead := ioutil.ReadFile(tools.PathKeyStore + "/" + ac.Accounts[i].KeyStoreFile)
	if errRead != nil {
		return common.Address{}, nil, errRead
	}
	key, errDecrypt := keystore.DecryptKey(keyjson, ac.Accounts[i].Password)
	if errDecrypt != nil {
		return common.Address{}, nil, errDecrypt
	}
	log.WithFields(log.Fields{"private_key": hexutil.Encode(crypto.FromECDSA(key.PrivateKey))}).Debug("The private key of the wallet writer")
	return key.Address, key.PrivateKey, nil
}
