package routers 

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUsersRoutes(router)
	router = SetTaskRoutes(router)
	ruter = SetNotesRoutes(router)
	return router
}