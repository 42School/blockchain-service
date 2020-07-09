package models

import (
	"github.com/ethereum/go-ethereum/common"
	crypgo "github.com/ethereum/go-ethereum/crypto"
	account "github.com/lpieri/42-Diploma/src/account"
	"github.com/lpieri/42-Diploma/src/contracts"
	"github.com/lpieri/42-Diploma/src/global"
	"log"
	"time"
)

type Diploma struct {
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	BirthDate	time.Time	`json:"birth_date"`
	AlumniDate	time.Time	`json:"alumni_date"`
	Level		float64		`json:"level"`
	Skills		[]float64	`json:"skills"`
}

func convertSkillToInt(skills []float64) [30]uint64 {
	newSkills := [30]uint64{}
	for i := 0; i < 30; i++ {
		newSkills[i] = uint64(skills[i] * 100)
	}
	return newSkills
}

func convertSkillToFloat(skills [30]uint64) [30]float64 {
	newSkills := [30]float64{}
	for i := 0; i < 30; i++ {
		newSkills[i] = float64(skills[i]) / 100
	}
	return newSkills
}

func convertDpToData(_dp Diploma, _sign []byte, _hash common.Hash) (uint64, [30]uint64, uint8, [32]byte, [32]byte, [32]byte) {
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

func NewDiploma(new Diploma) bool {
	account.CreateAccountsManager() // à mettre dans le main ?!
	dataToHash := new.FirstName + ", " + new.LastName + ", " + new.BirthDate.String()[:10] + ", " + new.AlumniDate.String()[:10]
	newHash := crypgo.Keccak256Hash([]byte(dataToHash))
	sign, err := account.KeyStore.SignHashWithPassphrase(account.GetAccount(), global.PasswordAccount, newHash.Bytes())
	if err != nil {
		return false
	}
	if contracts.CallCreateDiploma(convertDpToData(new, sign, newHash)) == false {
		// mettre le diplome dans la queue de retry !
		return false
	}
	return true
}

func GetDiploma(_dp Diploma) (float64, [30]float64, error) {
	dataToHash := _dp.FirstName + ", " + _dp.LastName + ", " + _dp.BirthDate.String()[:10] + ", " + _dp.AlumniDate.String()[:10]
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

func CheckDiploma(dp Diploma) bool {
	if dp.FirstName == "" || dp.LastName == "" || dp.Level <= 6 || len(dp.Skills) != 30 || dp.AlumniDate.IsZero() || dp.BirthDate.IsZero() {
		return false
	}
	for i := 0; i < len(dp.Skills); i++ {
		if dp.Skills[i] < 0.0 {
			return false
		}
	}
	return true
}

func PrintDiploma(dp Diploma) {
	log.Print("Enter in NewDiploma")
	log.Println("First Name:", dp.FirstName)
	log.Println("Last Name:", dp.LastName)
	log.Println("Birth Date:", dp.BirthDate)
	log.Println("Alumni Date:", dp.AlumniDate)
	log.Println("Level:", dp.Level)
	log.Println("Skills:", dp.Skills)
}
