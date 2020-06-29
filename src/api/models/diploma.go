package models

import (
	"log"
	"os"
	"time"
	account "github.com/lpieri/42-Diploma/src/account"
	crypgo "github.com/ethereum/go-ethereum/crypto"
)

type Diploma struct {
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	BirthDate	time.Time	`json:"birth_date"`
	AlumniDate	time.Time	`json:"alumni_date"`
	Level		float64		`json:"level"`
	Skills		[]float64	`json:"skills"`
}

var passwordAccount string = os.Getenv("KEYPASSWD")

func NewDiploma(new Diploma) bool {
	PrintDiploma(new)
	log.Println("Enter in NewDiploma")
	account.CreateAccountsManager()
	dataToHash := new.FirstName + ", " + new.LastName + ", " + new.BirthDate.String()[:10] + ", " + new.AlumniDate.String()[:10]
	newHash := crypgo.Keccak256Hash([]byte(dataToHash))
	log.Println(newHash.Hex())
	sign, err := account.KeyStore.SignHashWithPassphrase(account.GetAccount(), passwordAccount, newHash.Bytes())
	if err != nil {
		return false
	}
	log.Println(sign, err)
	return true
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
