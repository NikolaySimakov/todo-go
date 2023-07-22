package routes

import (
	"github.com/NikolaySimakov/todo-go/controllers/basic"
	"github.com/NikolaySimakov/todo-go/controllers/stack"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	// Basic TODO
	route.HandleFunc("/", basic.Show).Methods("GET")
	route.HandleFunc("/", basic.Add).Methods("POST")
	route.HandleFunc("/{id}", basic.Complete).Methods("PUT")
	route.HandleFunc("/{id}", basic.Delete).Methods("DELETE")

	// Stack TODO
	route.HandleFunc("/stack", stack.Current).Methods("GET")
	route.HandleFunc("/stack", stack.Add).Methods("POST")
	route.HandleFunc("/stack", stack.Complete).Methods("DELETE")

	return route
}
