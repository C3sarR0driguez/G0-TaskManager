package routers


import (
	"github.com/gorilla/mux"
	"githbub.com/C3sarR0driguez/TaskManager/controllers"
)

func SetUsersRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}

