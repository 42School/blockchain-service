package mocks

import (
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

type MockBlockchainImpl struct {
	mock.Mock
}

func (m MockBlockchainImpl) connectEthGetInstance() (*contracts.Diploma, *ethclient.Client, error) {
	args := m.Called()
	return args.Get(0).(*contracts.Diploma), args.Get(1).(*ethclient.Client), args.Error(2)
}

func (m MockBlockchainImpl) getAuth() (*bind.TransactOpts, error) {
	args := m.Called()
	return args.Get(0).(*bind.TransactOpts), args.Error(1)
}

func (m MockBlockchainImpl) getLogs(client *ethclient.Client) ([]types.Log, abi.ABI, error) {
	args := m.Called()
	return args.Get(0).([]types.Log), args.Get(1).(abi.ABI), args.Error(2)
}

func (m MockBlockchainImpl) GetBalance(address common.Address) (int64, error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}

func (m MockBlockchainImpl) GetRevert(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) string {
	args := m.Called()
	return args.String(0)
}

func (m MockBlockchainImpl) CheckSecurity(client *ethclient.Client, tx *types.Transaction, hash []byte) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m MockBlockchainImpl) CallCreateDiploma(level uint64, skills [30]uint64, skillsSlugs [30]string, v uint8, r [32]byte, s [32]byte, hash [32]byte) (*types.Transaction, bool) {
	args := m.Called()
	return args.Get(0).(*types.Transaction), args.Bool(1)
}

func (m MockBlockchainImpl) CallGetDiploma(hash []byte) (uint64, []contracts.FtDiplomaSkill, error) {
	args := m.Called()
	return uint64(args.Int(0)), args.Get(1).([]contracts.FtDiplomaSkill), args.Error(2)
}

func (m MockBlockchainImpl) CallGetAllDiploma() ([]contracts.FtDiplomaDiplomas, error) {
	args := m.Called()
	return args.Get(0).([]contracts.FtDiplomaDiplomas), args.Error(1)
}
