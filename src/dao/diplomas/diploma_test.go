package diplomas

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var diploma = Diploma{uuid.UUID{0}, "Louise", "Pieri", "27-12-1998", "01-01-2021", 21.42, []Skill{
	{"Security",16.42},
	{"Unix",13.87},
	{"Adaptation & creativity",12.7},
	{"Company experience",11.22},
	{"Algorithms & AI",10.38},
	{"Group & interpersonal",10.13},
	{"Graphics",7.49},
	{"Rigor",6.6},
	{"Imperative programming",5.34},
	{"Technology integration",5.26},
	{"Web",5.2},
	{"Organization",5.04},
	{"Network & system administration",4.5},
	{"DB & Data",4.28},
	{"Object-oriented programming",4.2}}}

func TestDiploma_String(t *testing.T) {
	a := assert.New(t)
	str := "Louise, Pieri, 27-12-1998, 01-01-2021"
	a.Equal(str, diploma.String(), "Function Diploma.String are not valid.")
}

func TestDiploma_CheckDiploma(t *testing.T) {
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
	diplomaNotValid.Skills = []Skill{}
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (no skills).")
	// Check a invalid diploma ([1]skills{level = 0.0})
	diplomaNotValid.Skills = []Skill{{"Web",0.0}}
	a.Equal(false, diplomaNotValid.CheckDiploma(), "Function Diploma.CheckDiploma are not valid with not valid diploma (no skills).")
}

func TestDiploma_LogFields(t *testing.T) {
	a := assert.New(t)
	field := log.Fields{"first_name": "Louise", "last_name": "Pieri", "birth_date": "27-12-1998", "alumni_date": "01-01-2021"}
	a.Equal(field, diploma.LogFields(), "Function Diploma.LogFields are not valid.")
}