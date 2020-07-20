package api

import (
	"github.com/gorilla/mux"
	controllers "github.com/42School/blockchain-service/src/api/controllers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/create-diploma").Name("Create").HandlerFunc(controllers.CreateDiploma)
	router.Methods("POST").Path("/get-diploma").Name("Get").HandlerFunc(controllers.GetDiploma)
	return router
}
