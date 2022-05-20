package account

import (
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ac_1 = Account{"UTC--2020-07-24T08-25-31.194883000Z--aec7bdfb241e56c04acf5e1a2a49f147867b85b7", "password"}
var ac_2 = Account{"UTC--2020-07-24T08-24-31.985849000Z--fe5ac6a7bb66da6916becb74a4a3e00074cd2599", "password"}
var ac_3 = Account{"UTC--2020-07-24T08-19-17.983576000Z--cac03bac6965e6d8ca96537a0344cc506b32c2c7", "password"}
var account_test = AccountsImpl{KeyStore: keystore.NewKeyStore("../../config-dev/keystore-sign", keystore.StandardScryptN, keystore.StandardScryptP), CurrentAccount: 0, Accounts: []Account{ac_1, ac_2, ac_3}}

func TestAccountsImpl_GetLenAccounts(t *testing.T) {
	a := assert.New(t)
	length := account_test.GetLenAccounts()
	a.Equal(3, length, "The Accounts length doesn't equal to 3")
}

func TestAccountsImpl_ChangeWriter(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, account_test.CurrentAccount, "Value of Current Account at start")
	account_test.ChangeWriter()
	a.Equal(1, account_test.CurrentAccount, "Value of Current Account doesn't equal at 1")
	account_test.ChangeWriter()
	a.Equal(2, account_test.CurrentAccount, "Value of Current Account doesn't equal at 2")
	account_test.ChangeWriter()
	a.Equal(0, account_test.CurrentAccount, "Value of Current Account doesn't equal at 0")
	account_test.ChangeWriter()
	a.Equal(1, account_test.CurrentAccount, "Value of Current Account doesn't equal at 1")
}

func TestAccountsImpl_GetSign(t *testing.T) {
	a := assert.New(t)
	expected := accounts.Account{common.HexToAddress("0x7e12234e994384a757e2689addb2a463ccd3b47d"), accounts.URL{Scheme:"keystore", Path:"/Users/louise/l0/blockchain-service/config-dev/keystore-sign/UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d"}}
	ac := account_test.GetSign()
	a.Equal(expected, ac)
}

func TestAccountsImpl_SignHash(t *testing.T) {
	a := assert.New(t)
	tools.PasswordAccount = "password"
	expected := []byte{40, 105, 34, 26, 74, 99, 82, 167, 136, 220, 171, 65, 147, 1, 129, 48, 213, 0, 138, 83, 165, 153, 50, 85, 235, 246, 224, 216, 122, 22, 167, 112, 77, 78, 1, 26, 74, 10, 174, 110, 141, 138, 218, 169, 46, 212, 158, 209, 58, 226, 20, 160, 185, 22, 86, 159, 144, 164, 160, 85, 161, 152, 72, 127, 0}
	sign, _ := account_test.SignHash(common.HexToHash("0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b"))
	a.Equal(expected, sign)
	tools.PasswordAccount = ""
	sign, _ = account_test.SignHash(common.HexToHash("0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b"))
	a.Equal([]byte(nil), sign)
}

func TestAccountsImpl_GetWriterByI(t *testing.T) {
	tools.PathKeyStore = "../../config-dev/keystore"
	a := assert.New(t)
	address, _, error := account_test.GetWriterByI(0)
	a.Equal(common.HexToAddress("0xaEc7BdfB241e56C04acF5E1a2a49F147867B85B7"), address, "")
	a.Equal(nil, error)
	address, _, error = account_test.GetWriterByI(1)
	a.Equal(common.HexToAddress("0xFE5Ac6A7Bb66Da6916Becb74A4A3e00074cd2599"), address, "")
	a.Equal(nil, error)
	address, _, error = account_test.GetWriterByI(2)
	a.Equal(common.HexToAddress("0xcac03bac6965e6d8ca96537a0344cc506b32c2c7"), address, "")
	a.Equal(nil, error)
}

func TestAccountsImpl_GetWriter(t *testing.T) {
	tools.PathKeyStore = "../../config-dev/keystore"
	a := assert.New(t)
	account_test.CurrentAccount = 1
	address, _, error := account_test.GetWriter()
	a.Equal(common.HexToAddress("0xFE5Ac6A7Bb66Da6916Becb74A4A3e00074cd2599"), address, "")
	account_test.CurrentAccount = 0
	address, _, error = account_test.GetWriter()
	a.Equal(common.HexToAddress("0xaEc7BdfB241e56C04acF5E1a2a49F147867B85B7"), address, "")
	account_test.CurrentAccount = 2
	address, _, error = account_test.GetWriter()
	a.Equal(common.HexToAddress("0xcac03bac6965e6d8ca96537a0344cc506b32c2c7"), address, "")
	a.Equal(nil, error)
}