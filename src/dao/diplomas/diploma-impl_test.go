package diplomas

import (
	"errors"
	"fmt"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

type SuiteEthWriting struct {
	suite.Suite
	hash        [32]byte
	v           uint8
	s           [32]byte
	r           [32]byte
	level       uint64
	skills      [30]uint64
	skillsSlugs [30]string
}

var diploma = DiplomaImpl{Id: uuid.UUID{0}, FirstName: "Louise", LastName: "Pieri", BirthDate: "1998-12-27", AlumniDate: "2021-01-01", Level: 21.42,
	Skills: []api.Skill{{"Security", 16.42}, {"Unix", 13.87}, {"Adaptation & creativity", 12.7}, {"Company experience", 11.22},
		{"Algorithms & AI", 10.38}, {"Group & interpersonal", 10.13}, {"Graphics", 7.49}, {"Rigor", 6.6},
		{"Imperative programming", 5.34}, {"Technology integration", 5.26}, {"Web", 5.2},
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

func TestDiplomaImpl_EthWriting(t *testing.T) {
	suite.Run(t, new(SuiteEthWriting))
}

func (s *SuiteEthWriting) SetupSuite() {
	hashByte, _ := hexutil.Decode("0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b")
	s.v = 27
	s.r = [32]byte{40, 105, 34, 26, 74, 99, 82, 167, 136, 220, 171, 65, 147, 1, 129, 48, 213, 0, 138, 83, 165, 153, 50, 85, 235, 246, 224, 216, 122, 22, 167, 112}
	s.s = [32]byte{77, 78, 1, 26, 74, 10, 174, 110, 141, 138, 218, 169, 46, 212, 158, 209, 58, 226, 20, 160, 185, 22, 86, 159, 144, 164, 160, 85, 161, 152, 72, 127}
	copy(s.hash[:], hashByte)
	s.level = 2142
	s.skills = [30]uint64{1642, 1387, 127, 1122, 1038, 1013, 749, 66, 534, 526, 52, 504, 45, 428, 42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	s.skillsSlugs = [30]string{
		"Security", "Unix", "Adaptation & creativity", "Company experience", "Algorithms & AI", "Group & interpersonal", "Graphics", "Rigor", "Imperative programming",
		"Technology integration", "Web", "Organization", "Network & system administration", "DB & Data", "Object-oriented programming", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	}
}

func (s *SuiteEthWriting) Test_No_Error() {
	mockBc := &contracts.MockBlockchainImpl{}
	mockBc.On("CallCreateDiploma", s.level, s.skills, s.skillsSlugs, s.v, s.r, s.s, s.hash).Return(&types.Transactions{}, true)
	diploma.blockchain = mockBc
	_, val := diploma.EthWriting()
	s.Equal(true, val)
}

func TestDiplomaImpl_EthGetter(t *testing.T) {
	suite.Run(t, new(SuiteEthGetter))
}

func (s *SuiteEthGetter) SetupSuite() {
	s.hash = "0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b"
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

func (s *SuiteEthGetter) Test_No_Error() {
	skillsBC := []contracts.FtDiplomaSkill{{"Security", 1642}, {"Unix", 1387}, {"Adaptation & creativity", 127}, {"Company experience", 1122},
		{"Algorithms & AI", 1038}, {"Group & interpersonal", 1013}, {"Graphics", 749}, {"Rigor", 66}, {"Imperative programming", 534},
		{"Technology integration", 526}, {"Web", 52}, {"Organization", 504}, {"Network & system administration", 45}, {"DB & Data", 428},
		{"Object-oriented programming", 42}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}}
	mockBc := &contracts.MockBlockchainImpl{}
	hashByte, _ := hexutil.Decode(s.hash)
	mockBc.On("CallGetDiploma", hashByte).Return(2121, skillsBC, nil)
	diploma.blockchain = mockBc
	level, skills, _ := diploma.EthGetter()
	s.Equal(21.21, level)
	s.Equal(len(skillsBC), len(skills))
	for i := 0; i < len(skills); i++ {
		s.Equal(skillsBC[i].Name, skills[i].Name)
		s.Equal(float64(skillsBC[i].Level)/100, skills[i].Level)
	}
}
