package handlers

import (
	"encoding/json"
	"errors"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SuiteGetDpHandler struct {
	suite.Suite
}

func (s *SuiteGetDpHandler) SetupSuite() {}

func TestGetDiplomaHandler_ServeHTTP(t *testing.T) {
	suite.Run(t, new(SuiteGetDpHandler))
}

func (s *SuiteGetDpHandler) Test_Error_JSON() {
	expectedString := "Fail Unmarshalling json\n"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("EthGetter").Return(0, [30]api.Skill{}, nil)
	diploma.On("ReadJson").Return(diploma, errors.New("Error"))
	u := &GetDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteGetDpHandler) Test_Error_Read_BC() {
	expectedString := "The request is fail, please retry or check the data.\n"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("EthGetter").Return(0, [30]api.Skill{}, errors.New("Error"))
	diploma.On("ReadJson").Return(diploma, nil)
	u := &GetDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteGetDpHandler) Test_No_Error() {
	skills := [30]api.Skill{
		{"Security",16.42}, {"Unix",13.87}, {"Adaptation & creativity",12.7},
		{"Company experience",11.22}, {"Algorithms & AI",10.38}, {"Group & interpersonal",10.13},
		{"Graphics",7.49}, {"Rigor",6.6}, {"Imperative programming",5.34},
		{"Technology integration",5.26}, {"Web",5.2}, {"Organization",5.04},
		{"Network & system administration",4.5}, {"DB & Data",4.28}, {"Object-oriented programming",4.2},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
	}
	expectedData := ResponseJson{true, "", ResponseData{"", 21, skills[:]}}
	expectedString, _ := json.Marshal(expectedData)
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("EthGetter").Return(21, skills, nil)
	diploma.On("ReadJson").Return(diploma, nil)
	u := &GetDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(string(expectedString), w.Body.String())
}