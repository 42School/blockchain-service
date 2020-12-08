package api

import (
	controllers "github.com/42School/blockchain-service/src/api/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/create-diploma").Name("Create").HandlerFunc(controllers.CreateDiploma)
	router.Methods("POST").Path("/get-diploma").Name("Get").HandlerFunc(controllers.GetDiploma)
	// Testing Route
	router.Methods("POST").Path("/check-request").Name("Check").HandlerFunc(controllers.CheckRouter)
	return router
}
