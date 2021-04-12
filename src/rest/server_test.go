package rest

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
)

type ServerMock struct {
	mock.Mock
	router *mux.Router
	Server *httptest.Server
}

func (m *ServerMock) initRouter() {
	m.router = mux.NewRouter().StrictSlash(true)
	m.router.Methods("POST").Path("/create-diploma")
	m.router.Methods("POST").Path("/get-diploma")
	m.router.Methods("GET").Path("/get-all-diploma")
}

func NewServerMock() *ServerMock {
	mock := &ServerMock{}
	mock.initRouter()
	mock.Server = httptest.NewServer(mock.router)
	return mock
}