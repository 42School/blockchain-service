package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

type MockBlockchainImpl struct {
	mock.Mock
}

func (m MockBlockchainImpl) GetBalance(address common.Address) (int64, error) {
	args := m.Called(address)
	return int64(args.Int(0)), args.Error(1)
}

func (m MockBlockchainImpl) GetRevert(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) string {
	args := m.Called(client, tx, receipt)
	return args.String(0)
}

func (m MockBlockchainImpl) CheckSecurity(client *ethclient.Client, tx *types.Transaction, hash []byte) bool {
	args := m.Called(client, tx, hash)
	return args.Bool(0)
}

func (m MockBlockchainImpl) CallCreateDiploma(level uint64, skills [30]uint64, skillsSlugs [30]string, v uint8, r [32]byte, s [32]byte, hash [32]byte) (*types.Transaction, bool) {
	args := m.Called(level, skills, skillsSlugs, v, r, s, hash)
	return args.Get(0).(*types.Transaction), args.Bool(1)
}

func (m MockBlockchainImpl) CallGetDiploma(hash []byte) (uint64, []FtDiplomaSkill, error) {
	args := m.Called(hash)
	return uint64(args.Int(0)), args.Get(1).([]FtDiplomaSkill), args.Error(2)
}

func (m MockBlockchainImpl) CallGetAllDiploma() ([]FtDiplomaDiplomas, error) {
	args := m.Called()
	return args.Get(0).([]FtDiplomaDiplomas), args.Error(1)
}
