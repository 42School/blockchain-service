package diplomas

import (
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common"
	crypgo "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"time"
)

func (_dp Diploma) convertDpToData(_sign []byte, _hash common.Hash) (uint64, [30]uint64, [30]string, uint8, [32]byte, [32]byte, [32]byte) {
	level := uint64(_dp.Level * 100)
	skillsLevels := [30]uint64{}
	skillsSlugs := [30]string{}
	for i := 0; i < 30; i++ {
		if i > len(_dp.Skills)-1 {
			skillsLevels[i] = uint64(0)
			skillsSlugs[i] = ""
		} else {
			skillsLevels[i] = uint64(_dp.Skills[i].Level * 100)
			skillsSlugs[i] = _dp.Skills[i].Name
		}
	}
	v := uint8(int(_sign[64])) + 27
	r := [32]byte{}
	s := [32]byte{}
	hash := [32]byte{}
	copy(r[:], _sign[:32])
	copy(s[:], _sign[32:64])
	copy(hash[:], _hash.Bytes())
	return level, skillsLevels, skillsSlugs, v, r, s, hash
}

func (_dp Diploma) EthWriting() (string, bool) {
	dataToHash := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	newHash := crypgo.Keccak256Hash([]byte(dataToHash))
	sign, err := account.KeyStore.SignHashWithPassphrase(account.GetSignAccount(), tools.PasswordAccount, newHash.Bytes())
	log.WithFields(log.Fields{"hash": newHash.String(), "sign": common.Bytes2Hex(sign)}).Debug("The hash & signature of the diploma")
	if err != nil {
		tools.LogsError(err)
		return "", false
	}
	tx, success := contracts.CallCreateDiploma(_dp.convertDpToData(sign, newHash))
	if success == false {
		_dp.AddToRetry()
		return "", false
	}
	log.WithFields(log.Fields{"hash": newHash.String(), "tx": tx.Hash().String()}).Info("Diploma submit in transaction.")
	addToCheck(VerificationHash{Tx: tx, StudentHash: newHash.Bytes(), SendTime: time.Now()})
	return newHash.Hex(), true
}

func (_dp Diploma) EthGetter() (float64, [30]Skill, error) {
	dataToHash := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	hash := crypgo.Keccak256Hash([]byte(dataToHash))
	levelInt, skillsEth, err := contracts.CallGetDiploma(hash.Bytes())
	if err != nil {
		tools.LogsError(err)
		return 0, [30]Skill{}, err
	}
	level := float64(levelInt) / 100
	skills := [30]Skill{}
	for i := 0; i < 30; i++ {
		skills[i].Level = float64(skillsEth[i].Level) / 100
		skills[i].Name = skillsEth[i].Slug
	}
	log.Print(level, skills)
	return level, skills, nil
}

func EthAllGetter() []contracts.FtDiplomaDiplomas {
	diplomas, err := contracts.CallGetAllDiploma()
	if err != nil {
		tools.LogsError(err)
		return nil
	}
	return diplomas
}
