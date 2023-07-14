package main

import (
	"log"
	"net/http"

	"github.com/NikolaySimakov/todo-go/routes"
)

func main() {
	err := http.ListenAndServe(":"+"3000", routes.Init())
	if err != nil {
		log.Fatal(err)
	}
}
