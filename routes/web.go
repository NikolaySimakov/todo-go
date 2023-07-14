package routes

import (
	"github.com/NikolaySimakov/todo-go/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)
	return route
}
