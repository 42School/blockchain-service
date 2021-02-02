package diplomas

import (
	"context"
	account "github.com/42School/blockchain-service/src/account"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/42School/blockchain-service/src/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	crypgo "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Diploma struct {
	Id			uuid.UUID   `bson:"_id"`
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	BirthDate	string		`json:"birth_date"`
	AlumniDate	string		`json:"alumni_date"`
	Level		float64		`json:"level"`
	Skills		[]float64	`json:"skills"`
}

type VerificationHash struct {
	Id	uuid.UUID   `bson:"_id"`
	Tx *types.Transaction
	StudentHash []byte
}

func convertSkillToInt(skills []float64) [30]uint64 {
	newSkills := [30]uint64{}
	for i := 0; i < 30; i++ {
		if i > len(skills) - 1 {
			newSkills[i] = uint64(0)
		} else {
			newSkills[i] = uint64(skills[i] * 100)
		}
	}
	return newSkills
}

func convertSkillToFloat(skills [30]uint64) [30]float64 {
	newSkills := [30]float64{}
	for i := 0; i < 30; i++ {
		if i > len(skills) - 1 {
			newSkills[i] = float64(0)
		} else {
			newSkills[i] = float64(skills[i]) / 100
		}
	}
	return newSkills
}

func addToCheck(toAdd VerificationHash) {
	var checkDB VerificationHash
	result := tools.ToCheckDB.FindOne(context.TODO(), bson.M{"studenthash": toAdd.StudentHash})
	err := result.Decode(&checkDB)
	if hexutil.Encode(checkDB.StudentHash) == hexutil.Encode(toAdd.StudentHash) || err == nil {
		tools.LogsDev("Verification Hash already in DB Queue: " + hexutil.Encode(toAdd.StudentHash))
		return
	}
	tools.ToCheckHash.PushBack(toAdd)
	txJson, _ := toAdd.Tx.MarshalJSON()
	tools.ToCheckDB.InsertOne(context.Background(), bson.M{"tx": txJson, "studenthash": toAdd.StudentHash})
}

func (_dp Diploma) CheckDiploma() bool {
	if _dp.FirstName == "" || _dp.LastName == "" || _dp.Level <= 6 || _dp.AlumniDate == "" || _dp.BirthDate == "" {
		return false
	}
	for i := 0; i < len(_dp.Skills); i++ {
		if _dp.Skills[i] <= 0.0 {
			return false
		}
	}
	return true
}

func (_dp Diploma) PrintDiploma() {
	log.Println("First Name:", _dp.FirstName)
	log.Println("Last Name:", _dp.LastName)
	log.Println("Birth Date:", _dp.BirthDate)
	log.Println("Alumni Date:", _dp.AlumniDate)
	log.Println("Level:", _dp.Level)
	log.Println("Skills:", _dp.Skills)
}

func (_dp Diploma) String() string {
	str := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	return str
}

func (_dp Diploma) AddToRetry() {
	copyList := tools.RetryQueue
	for e := copyList.Front(); e != nil; e = e.Next() {
		if e != nil {
			diploma, _ := e.Value.(Diploma)
			tools.LogsDev("Diploma in list: " + diploma.String())
			tools.LogsDev("Diploma to find: " + _dp.String())
			if diploma.String() == _dp.String() {
				tools.LogsDev("Match diploma in list & to find")
				return
			}
		}
	}
	var DpDB Diploma
	result := tools.RetryDB.FindOne(context.TODO(), bson.M{"firstname": _dp.FirstName, "lastname": _dp.LastName, "birthdate": _dp.BirthDate, "alumnidate": _dp.AlumniDate})
	err := result.Decode(&DpDB)
	if DpDB.String() == _dp.String() || err == nil {
		tools.LogsDev("Diploma already in DB Queue: " + _dp.String())
		return
	}
	tools.LogsDev("Adding diploma in Queue: " + _dp.String())
	_dp.Id = uuid.New()
	tools.RetryDB.InsertOne(context.TODO(), _dp)
	tools.RetryQueue.PushBack(_dp)
}

func (_dp Diploma) convertDpToData(_sign []byte, _hash common.Hash) (uint64, [30]uint64, uint8, [32]byte, [32]byte, [32]byte) {
	level := uint64(_dp.Level * 100)
	skills := convertSkillToInt(_dp.Skills)
	v := uint8(int(_sign[64])) + 27
	r := [32]byte{}
	s := [32]byte{}
	hash := [32]byte{}
	copy(r[:], _sign[:32])
	copy(s[:], _sign[32:64])
	copy(hash[:], _hash.Bytes())
	return level, skills, v, r, s, hash
}

func (_dp Diploma) EthWriting() (string, bool) {
	dataToHash := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	newHash := crypgo.Keccak256Hash([]byte(dataToHash))
	sign, err := account.KeyStore.SignHashWithPassphrase(account.GetSignAccount(), tools.PasswordAccount, newHash.Bytes())
	tools.LogsDev("The hash of the diploma is " + newHash.String())
	tools.LogsDev("The signature on the diploma is " + common.Bytes2Hex(sign))
	if err != nil {
		return "", false
	}
	tx, success := contracts.CallCreateDiploma(_dp.convertDpToData(sign, newHash))
	if success == false {
		_dp.AddToRetry()
		return "", false
	}
	addToCheck(VerificationHash{Tx: tx, StudentHash: newHash.Bytes()})
	return newHash.Hex(), true
}

func (_dp Diploma) EthGetter() (float64, [30]float64, error) {
	dataToHash := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate + ", " + _dp.AlumniDate
	hash := crypgo.Keccak256Hash([]byte(dataToHash))
	levelInt, skillsInt, err := contracts.CallGetDiploma(hash.Bytes())
	if err != nil {
		return 0, [30]float64{}, err
	}
	level := float64(levelInt) / 100
	skills := convertSkillToFloat(skillsInt)
	log.Print(levelInt, skillsInt)
	return level, skills, nil
}

func EthAllGetter() []contracts.FtDiplomaDiploma {
	diplomas, err := contracts.CallGetAllDiploma()
	if err != nil {
		return nil
	}
	return diplomas
}