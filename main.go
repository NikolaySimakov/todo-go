package main

import (
	"log"
	"net/http"
	"time"

	"github.com/NikolaySimakov/todo-go/routes"
)

func main() {
	srv := &http.Server{
		Handler:      routes.Init(),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server will start at http://%s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
