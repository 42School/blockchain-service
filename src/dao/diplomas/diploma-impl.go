package diplomas

import (
	"encoding/json"
	"github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	crypgo "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"time"
)

type WebhookData struct {
	Login               string `json:"login"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	BirthDate           string `json:"birth_date"`
	AlumnizedCursusUser int    `json:"alumnized_cursus_user"`
}

type DiplomaImpl struct {
	Id         uuid.UUID   `bson:"_id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	BirthDate  string      `json:"birth_date"`
	AlumniDate string      `json:"alumni_date"`
	Level      float64     `json:"level"`
	Skills     []api.Skill `json:"skills"`
	Counter    int         `json:"counter"`
}

func (_dp DiplomaImpl) ReadWebhook(body io.ReadCloser) (Diploma, error) {
	var webhookData WebhookData
	err := json.NewDecoder(body).Decode(&webhookData)
	if err != nil {
		tools.LogsError(err)
		return _dp, err
	}
	level, skills, err := api.GetCursusUser(webhookData.AlumnizedCursusUser)
	if err != nil {
		tools.LogsError(err)
		return _dp, err
	}
	_dp.FirstName = webhookData.FirstName
	_dp.LastName = webhookData.LastName
	_dp.BirthDate = webhookData.BirthDate
	_dp.AlumniDate = "2021-01-01"
	//_dp.AlumniDate = time.Now().Format("2006-01-02")
	_dp.Level = level
	_dp.Skills = skills
	_dp.Counter = 0
	log.WithFields(_dp.LogFields()).Debug("Webhook to Diploma success")
	return _dp, nil
}

func (_dp DiplomaImpl) ReadJson(body io.ReadCloser) (Diploma, error) {
	err := json.NewDecoder(body).Decode(&_dp)
	if err != nil {
		return _dp, err
	}
	return _dp, nil
}

func addToCheck(toAdd VerificationHash) {
	var checkDB VerificationHash
	result := tools.Db.FindOneCheck(bson.M{"studenthash": toAdd.StudentHash})
	err := result.Decode(&checkDB)
	if hexutil.Encode(checkDB.StudentHash) == hexutil.Encode(toAdd.StudentHash) || err == nil {
		log.WithFields(log.Fields{"hash": hexutil.Encode(toAdd.StudentHash)}).Debug("The verification hash already exist in the CheckDB")
		return
	}
	tools.ToCheckHash.PushBack(toAdd)
	txJson, _ := toAdd.Tx.MarshalJSON()
	metrics.GaugeCheckQueue.Inc()
	metrics.CounterCheckQueue.Inc()
	tools.Db.InsertOneCheck(bson.M{"tx": txJson, "studenthash": toAdd.StudentHash, "time": toAdd.SendTime})
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
	result := tools.Db.FindOneRetry(bson.M{"firstname": _dp.FirstName, "lastname": _dp.LastName, "birthdate": _dp.BirthDate, "alumnidate": _dp.AlumniDate})
	err := result.Decode(&DpDB)
	if DpDB.String() == _dp.String() || err == nil {
		log.WithFields(_dp.LogFields()).Debug("Diploma already exist in Retry DB Queue")
		return
	}
	_dp.Id = uuid.New()
	if _dp.Counter == 0 {
		_dp.Counter = 1
	}
	tools.Db.InsertOneRetry(_dp)
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
	hash := crypgo.Keccak256Hash([]byte(_dp.String()))
	sign, err := account.Accounts.SignHash(hash)
	log.WithFields(log.Fields{"hash": hash.String(), "sign": common.Bytes2Hex(sign)}).Debug("The hash & signature of the diploma")
	if err != nil {
		tools.LogsError(err)
		return "", false
	}
	tx, success := contracts.Blockchain.CallCreateDiploma(_dp.convertDpToData(sign, hash))
	//tx, success := _dp.blockchain.CallCreateDiploma(_dp.convertDpToData(sign, hash))
	if success == false {
		_dp.AddToRetry()
		return "", false
	}
	metrics.NumberOfRetryDiploma.Observe(float64(_dp.Counter))
	metrics.NumberOfRetryPerDiploma.WithLabelValues(_dp.String()).Add(float64(_dp.Counter))
	log.WithFields(log.Fields{"hash": hash.String(), "tx": tx.Hash().String()}).Info("Diploma submit in transaction.")
	addToCheck(VerificationHash{Tx: tx, StudentHash: hash.Bytes(), SendTime: time.Now()})
	return hash.Hex(), true
}

func (_dp DiplomaImpl) EthGetter() (float64, [30]api.Skill, error) {
	hash := crypgo.Keccak256Hash([]byte(_dp.String()))
	//log.Debug(hash.String())
	levelInt, skillsEth, err := contracts.Blockchain.CallGetDiploma(hash.Bytes())
	if err != nil {
		tools.LogsError(err)
		return 0, [30]api.Skill{}, err
	}
	level := float64(levelInt) / 100
	skills := [30]api.Skill{}
	for i := 0; i < 30; i++ {
		skills[i].Level = float64(skillsEth[i].Level) / 100
		skills[i].Name = skillsEth[i].Name
	}
	log.Debug(level, skills)
	return level, skills, nil
}
