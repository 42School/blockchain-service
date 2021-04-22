package diplomas

import (
	"errors"
	"fmt"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteEthGetter struct {
	suite.Suite
	hash string
}

var diploma = DiplomaImpl{Id: uuid.UUID{0}, FirstName: "Louise", LastName: "Pieri", BirthDate: "1998-12-27", AlumniDate: "2021-01-01", Level: 21.42,
	Skills: []api.Skill{{"Security", 16.42}, {"Unix", 13.87}, {"Adaptation & creativity", 12.7}, {"Company experience", 11.22},
	{"Algorithms & AI", 10.38}, {"Group & interpersonal", 10.13}, {"Graphics", 7.49}, {"Rigor", 6.6},
	{"Imperative programming", 5.34}, {"Technology integration",5.26}, {"Web",5.2},
	{"Organization", 5.04}, {"Network & system administration", 4.5}, {"DB & Data", 4.28}, {"Object-oriented programming", 4.2}}}

func TestDiplomaImpl_CheckDiploma(t *testing.T) {
	a := assert.New(t)
	// Check a 100% valid diploma
	a.Equal(true, diploma.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with valid diploma.")
	// Check a invalid diploma (first_name = "")
	diplomaNotValid := diploma
	diplomaNotValid.FirstName = ""
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (first_name = '').")
	// Check a invalid diploma (last_name = "")
	diplomaNotValid.FirstName = diploma.FirstName
	diplomaNotValid.LastName = ""
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (last_name = '').")
	// Check a invalid diploma (birth_date = "")
	diplomaNotValid.LastName = diploma.LastName
	diplomaNotValid.BirthDate = ""
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (birth_date = '').")
	// Check a invalid diploma (alumni_date = "")
	diplomaNotValid.BirthDate = diploma.BirthDate
	diplomaNotValid.AlumniDate = ""
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (alumni_date = '').")
	// Check a invalid diploma (level = 0.0)
	diplomaNotValid.AlumniDate = diploma.AlumniDate
	diplomaNotValid.Level = 0.0
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (level = 0.0).")
	// Check a invalid diploma (no skills)
	diplomaNotValid.Level = diploma.Level
	diplomaNotValid.Skills = []api.Skill{}
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (no skills).")
	// Check a invalid diploma ([1]skills{level = 0.0})
	diplomaNotValid.Skills = []api.Skill{{"Web", 0.0}}
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (no skills).")
}

func TestDiplomaImpl_String(t *testing.T) {
	a := assert.New(t)
	str := "Louise, Pieri, 1998-12-27, 2021-01-01"
	a.Equal(str, diploma.String(), "Function Diploma.String are not valid.")
}

func TestDiplomaImpl_LogFields(t *testing.T) {
	a := assert.New(t)
	field := log.Fields{"first_name": "Louise", "last_name": "Pieri", "birth_date": "1998-12-27", "alumni_date": "2021-01-01"}
	a.Equal(field, diploma.LogFields(), "Function Diploma.LogFields are not valid.")
}

func (s *SuiteEthGetter) SetupSuite() {
	s.hash = "0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b"
}

func TestDiplomaImpl_EthGetter(t *testing.T) {
	suite.Run(t, new(SuiteEthGetter))
}

func (s *SuiteEthGetter) Test_Error_Diploma_Doesnt_Exist() {
	mockBc := &contracts.MockBlockchainImpl{}
	hashByte, _ := hexutil.Decode(s.hash)
	mockBc.On("CallGetDiploma", hashByte).Return(0, []contracts.FtDiplomaSkill{}, errors.New("the diploma doesnt exist"))
	diploma.blockchain = mockBc
	level, _, err := diploma.EthGetter()
	s.Equal(fmt.Errorf("the diploma doesnt exist"), err)
	s.Equal(float64(0), level)
}