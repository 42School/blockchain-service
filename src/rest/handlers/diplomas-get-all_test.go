package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/contracts"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	skills := [30]contracts.FtDiplomaSkill{
		{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},
		{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},
		{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0},{"", 0}}
	hashByte, _ := hexutil.Decode("0xa41eeebbe22e2235a8ef94074c79c92ef6448baca12625ed6e26a61ddb60b55b")
	var hash [32]byte
	copy(hash[:], hashByte)
	sign := contracts.FtDiplomaSign{
		27,
		[32]byte{40, 105, 34, 26, 74, 99, 82, 167, 136, 220, 171, 65, 147, 1, 129, 48, 213, 0, 138, 83, 165, 153, 50, 85, 235, 246, 224, 216, 122, 22, 167, 112},
		[32]byte{77, 78, 1, 26, 74, 10, 174, 110, 141, 138, 218, 169, 46, 212, 158, 209, 58, 226, 20, 160, 185, 22, 86, 159, 144, 164, 160, 85, 161, 152, 72, 127}}
	oneData := contracts.FtDiplomaDiplomas{2142, skills, hash, sign}
	data := []contracts.FtDiplomaDiplomas{oneData, oneData}
	expectedString := "[{\"Level\":2142,\"Skills\":[{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0}],\"Hash\":[164,30,238,187,226,46,34,53,168,239,148,7,76,121,201,46,246,68,139,172,161,38,37,237,110,38,166,29,219,96,181,91],\"Signature\":{\"V\":27,\"R\":[40,105,34,26,74,99,82,167,136,220,171,65,147,1,129,48,213,0,138,83,165,153,50,85,235,246,224,216,122,22,167,112],\"S\":[77,78,1,26,74,10,174,110,141,138,218,169,46,212,158,209,58,226,20,160,185,22,86,159,144,164,160,85,161,152,72,127]}},{\"Level\":2142,\"Skills\":[{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0},{\"Name\":\"\",\"Level\":0}],\"Hash\":[164,30,238,187,226,46,34,53,168,239,148,7,76,121,201,46,246,68,139,172,161,38,37,237,110,38,166,29,219,96,181,91],\"Signature\":{\"V\":27,\"R\":[40,105,34,26,74,99,82,167,136,220,171,65,147,1,129,48,213,0,138,83,165,153,50,85,235,246,224,216,122,22,167,112],\"S\":[77,78,1,26,74,10,174,110,141,138,218,169,46,212,158,209,58,226,20,160,185,22,86,159,144,164,160,85,161,152,72,127]}}]\n"
	bc := contracts.MockBlockchainImpl{}
	bc.On("CallGetAllDiploma").Return(data, nil)
	u := &GetAllDiplomaHandler{bc, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-all-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	s.Equal(expectedString, w.Body.String())
}