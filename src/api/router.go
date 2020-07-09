package api

import (
	"github.com/gorilla/mux"
	controllers "github.com/lpieri/42-Diploma/src/api/controllers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/create-diploma").Name("Create").HandlerFunc(controllers.CreateDiploma)
	router.Methods("POST").Path("/get-diploma").Name("Get").HandlerFunc(controllers.GetDiploma)
	router.Methods("GET").Path("/check").Name("Check").HandlerFunc(controllers.CheckHash)
	return router
}
