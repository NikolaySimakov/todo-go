package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NikolaySimakov/todo-go/configs"
	"github.com/NikolaySimakov/todo-go/models"
	"github.com/gorilla/mux"
)

var (
	id        int
	title     string
	completed bool
	db        = configs.Database()
)

func Add(w http.ResponseWriter, r *http.Request) {

	// test task
	task := models.Task{
		Title:      "qwerty",
		IsComplete: false,
	}

	_, err := db.Exec("INSERT INTO tasks (title, is_completed) VALUES ($1, $2)", task.Title, task.IsComplete)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully inserted!")
}

func Show(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM tasks")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// creates tasks list
	var tasks models.List

	// run rows gererator
	for rows.Next() {
		err := rows.Scan(&id, &title, &completed)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, models.Task{
			Id:         id,
			Title:      title,
			IsComplete: completed,
		})
	}

	fmt.Println(tasks)
}

func Complete(w http.ResponseWriter, r *http.Request) {

	// getting id
	vars := mux.Vars(r)
	id := vars["id"]

	// set is_completed = true by task id
	_, err := db.Exec("UPDATE tasks SET is_completed = true WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	// getting id
	vars := mux.Vars(r)
	id := vars["id"]

	// delete row by id
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
