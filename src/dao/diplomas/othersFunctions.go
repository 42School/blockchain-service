package diplomas

import (
	"context"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (_dp Diploma) CheckDiploma() bool {
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

func (_dp Diploma) String() string {
	str := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	return str
}

func (_dp Diploma) LogFields() log.Fields {
	return log.Fields{"first_name": _dp.FirstName, "last_name": _dp.LastName, "birth_date": _dp.BirthDate, "alumni_date": _dp.AlumniDate}
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

func (_dp Diploma) AddToRetry() {
	copyList := tools.RetryQueue
	for e := copyList.Front(); e != nil; e = e.Next() {
		if e != nil {
			diploma, _ := e.Value.(Diploma)
			log.WithFields(diploma.LogFields()).Debug("Diploma in the retry queue")
			log.WithFields(_dp.LogFields()).Debug("Diploma to find in the retry queue")
			if diploma.String() == _dp.String() {
				log.Debug("Matching diploma in the queue & to find")
				return
			}
		}
	}
	var DpDB Diploma
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
