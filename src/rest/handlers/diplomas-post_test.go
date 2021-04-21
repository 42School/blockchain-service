package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDiplomaHandler_ServeHTTP(t *testing.T) {
	mockDiploma := mocks.MockDiplomaImpl{}
	mockDiploma.On("ReadWebhook").Return(mockDiploma, nil)
	mockDiploma.On("CheckDiploma").Return(false)
	mockDiploma.On("EthWriting").Return("0x9eeef4248948f749c71af80a097716decbb59d73d05666eed7642a7491107115", true)
	u := &CreateDiplomaHandler{mockDiploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/create-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
}
