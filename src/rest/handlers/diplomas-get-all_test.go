package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SuiteGetAllDpHandler struct {
	suite.Suite
}

func (s *SuiteGetAllDpHandler) SetupSuite() {}

func TestGetAllDiplomaHandler_ServeHTTP(t *testing.T) {
	suite.Run(t, new(SuiteGetAllDpHandler))
}

func (s *SuiteGetAllDpHandler) Test_Error_Recovery() {
	expectedString := "A problem occurred during data recovery.\n"
	bc := contracts.MockBlockchainImpl{}
	bc.On("CallGetAllDiploma").Return([]contracts.FtDiplomaDiplomas{}, errors.New("Error"))
	u := &GetAllDiplomaHandler{bc, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-all-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteGetAllDpHandler) Test_No_Error() {
	data := []contracts.FtDiplomaDiplomas{}
	//skills := [30]api.Skill{
	//	{"Security",16.42}, {"Unix",13.87}, {"Adaptation & creativity",12.7},
	//	{"Company experience",11.22}, {"Algorithms & AI",10.38}, {"Group & interpersonal",10.13},
	//	{"Graphics",7.49}, {"Rigor",6.6}, {"Imperative programming",5.34},
	//	{"Technology integration",5.26}, {"Web",5.2}, {"Organization",5.04},
	//	{"Network & system administration",4.5}, {"DB & Data",4.28}, {"Object-oriented programming",4.2},
	//	{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
	//}
	//expectedData := ResponseJson{true, "", ResponseData{"", 21, skills[:]}}
	//expectedString, _ := json.Marshal(...)
	bc := contracts.MockBlockchainImpl{}
	bc.On("CallGetAllDiploma").Return(data, nil)
	u := &GetAllDiplomaHandler{bc, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-all-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	//s.Equal(string(expectedString), w.Body.String())
}