package diplomas

import (
	"context"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	crypgo "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Skill struct {
	Name  string  `json:"name"`
	Level float64 `json:"level"`
}

type DiplomaImpl struct {
	Id         uuid.UUID `bson:"_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	AlumniDate string    `json:"alumni_date"`
	Level      float64   `json:"level"`
	Skills     []Skill   `json:"skills"`
	Counter    int       `json:"counter"`
}

func addToCheck(toAdd VerificationHash) {
	var checkDB VerificationHash
	result := tools.ToCheckDB.FindOne(context.TODO(), bson.M{"studenthash": toAdd.StudentHash})
	err := result.Decode(&checkDB)
	if hexutil.Encode(checkDB.StudentHash) == hexutil.Encode(toAdd.StudentHash) || err == nil {
		log.WithFields(log.Fields{"hash": hexutil.Encode(toAdd.StudentHash)}).Debug("The verification hash already exist in the CheckDB")
		return
	}
	tools.ToCheckHash.PushBack(toAdd)
	txJson, _ := toAdd.Tx.MarshalJSON()
	metrics.GaugeCheckQueue.Inc()
	metrics.CounterCheckQueue.Inc()
	tools.ToCheckDB.InsertOne(context.Background(), bson.M{"tx": txJson, "studenthash": toAdd.StudentHash, "time": toAdd.SendTime})
}

func (_dp DiplomaImpl) AddToRetry() {
	copyList := tools.RetryQueue
	for e := copyList.Front(); e != nil; e = e.Next() {
		if e != nil {
			diploma, _ := e.Value.(DiplomaImpl)
			log.WithFields(diploma.LogFields()).Debug("Diploma in the retry queue")
			log.WithFields(_dp.LogFields()).Debug("Diploma to find in the retry queue")
			if diploma.String() == _dp.String() {
				log.Debug("Matching diploma in the queue & to find")
				return
			}
		}
	}
	var DpDB DiplomaImpl
	result := tools.RetryDB.FindOne(context.TODO(), bson.M{"firstname": _dp.FirstName, "lastname": _dp.LastName, "birthdate": _dp.BirthDate, "alumnidate": _dp.AlumniDate})
	err := result.Decode(&DpDB)
	if DpDB.String() == _dp.String() || err == nil {
		log.WithFields(_dp.LogFields()).Debug("Diploma already exist in Retry DB Queue")
		return
	}
	_dp.Id = uuid.New()
	if _dp.Counter == 0 {
		_dp.Counter = 1
	}
	tools.RetryDB.InsertOne(context.TODO(), _dp)
	tools.RetryQueue.PushBack(_dp)
	metrics.GaugeRetryQueue.Inc()
	metrics.CounterRetryQueue.Inc()
	log.WithFields(_dp.LogFields()).Debug("Adding diploma in the retry queue & retry db")
}

func (_dp DiplomaImpl) CheckDiploma() bool {
	if _dp.FirstName == "" || _dp.LastName == "" || _dp.Level <= 6 || _dp.AlumniDate == "" || _dp.BirthDate == "" || len(_dp.Skills) == 0 {
		return false
	}
	for i := 0; i < len(_dp.Skills); i++ {
		if _dp.Skills[i].Level <= 0.0 {
			return false
		}
	}
	return true
}

func (_dp DiplomaImpl) String() string {
	str := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	return str
}

func (_dp DiplomaImpl) LogFields() log.Fields {
	return log.Fields{"first_name": _dp.FirstName, "last_name": _dp.LastName, "birth_date": _dp.BirthDate, "alumni_date": _dp.AlumniDate}
}

func (_dp DiplomaImpl) convertDpToData(_sign []byte, _hash common.Hash) (uint64, [30]uint64, [30]string, uint8, [32]byte, [32]byte, [32]byte) {
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

func (_dp DiplomaImpl) EthWriting() (string, bool) {
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
	metrics.NumberOfRetryDiploma.Observe(float64(_dp.Counter))
	metrics.NumberOfRetryPerDiploma.WithLabelValues(_dp.String()).Add(float64(_dp.Counter))
	log.WithFields(log.Fields{"hash": newHash.String(), "tx": tx.Hash().String()}).Info("Diploma submit in transaction.")
	addToCheck(VerificationHash{Tx: tx, StudentHash: newHash.Bytes(), SendTime: time.Now()})
	return newHash.Hex(), true
}

func (_dp DiplomaImpl) EthGetter() (float64, [30]Skill, error) {
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
		skills[i].Name = skillsEth[i].Name
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
