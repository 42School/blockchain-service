package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SuiteCreateDpHandler struct {
	suite.Suite
	hash    string
}

func (s *SuiteCreateDpHandler) SetupSuite() {
	s.hash = "0x9eeef4248948f749c71af80a097716decbb59d73d05666eed7642a7491107115"
}

func TestCreateDiplomaHandler_ServeHTTP(t *testing.T) {
	suite.Run(t, new(SuiteCreateDpHandler))
}

func (s *SuiteCreateDpHandler) Test_Error_JSON() {
	expectedString := "Fail Unmarshalling json\n"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("CheckDiploma").Return(true)
	diploma.On("EthWriting").Return(s.hash, true)
	diploma.On("ReadWebhook").Return(diploma, errors.New("Error"))
	u := &CreateDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/create-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteCreateDpHandler) Test_Error_CheckDiploma() {
	expectedString := "The data sent are not valid, to be written in blockchain please try again !\n"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("CheckDiploma").Return(false)
	diploma.On("EthWriting").Return(s.hash, true)
	diploma.On("ReadWebhook").Return(diploma, nil)
	u := &CreateDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/create-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteCreateDpHandler) Test_Error_Write_Diplomas() {
	expectedString := "Blockchain writing had a problem, the diploma is saved in the queue.\n"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("CheckDiploma").Return(true)
	diploma.On("EthWriting").Return("", false)
	diploma.On("ReadWebhook").Return(diploma, nil)
	u := &CreateDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/create-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}

func (s *SuiteCreateDpHandler) Test_No_Error() {
	expectedString := "{\"Status\":true,\"Message\":\"The writing in blockchain has been done, it will be confirmed in 10 min.\",\"Data\":{\"Hash\":\"" + s.hash + "\",\"Level\":0,\"Skills\":[]}}"
	diploma := mocks.MockDiplomaImpl{}
	diploma.On("CheckDiploma").Return(true)
	diploma.On("EthWriting").Return(s.hash, true)
	diploma.On("ReadWebhook").Return(diploma, nil)
	u := &CreateDiplomaHandler{diploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/create-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}
