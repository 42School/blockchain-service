package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/api"
	"github.com/42School/blockchain-service/src/dao/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDiplomaHandler_ServeHTTP(t *testing.T) {
	mockDiploma := mocks.MockDiplomaImpl{}
	mockDiploma.On("EthGetter").Return(0, [30]api.Skill{}, nil)
	mockDiploma.On("ReadJson").Return(mockDiploma, nil)
	u := &GetDiplomaHandler{mockDiploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
}
