package handlers

import (
	"errors"
	"github.com/42School/blockchain-service/src/dao/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDiplomaHandler_ServeHTTP(t *testing.T) {
	mockDiploma := mocks.MockDiplomaImpl{}
	mockDiploma.On("ReadJson").Return(nil)
	mockDiploma.On("EthGetter").Return(nil)
	u := &GetDiplomaHandler{mockDiploma, errors.New("")}
	r := httptest.NewRequest(http.MethodPost, "/get-diploma", nil)
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
}
