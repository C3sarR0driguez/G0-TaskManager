package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/C3sarR0driguez/TaskManager/common"
	"github.com/C3sarR0driguez/TaskManager/controllers"
)

func SetNotesRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
	noteRouter.HandleFunc("/notes/tasks/{id}", controllers.GeNodeByTask).Methods("GET")
	noteRouter.HandleFunc("/notes/{id}", controllers.DeleteNode).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter)
	))
	return router
}