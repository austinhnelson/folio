package routes

import (

	"github.com/gorilla/mux"
	"folio/handlers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/register", handlers.Register).Methods("POST")
	r.HandleFunc("/auth/Login", handlers.Login).Methods("POST")
	r.HandleFunc("/feed", handlers.GetFeed).Methods("GET")

	return r
}