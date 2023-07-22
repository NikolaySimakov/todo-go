package stack

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/NikolaySimakov/todo-go/configs"
	"github.com/NikolaySimakov/todo-go/models"
)

var (
	id        int
	title     string
	completed bool
	db        = configs.Database()
)

// show current (last) task
func Current(w http.ResponseWriter, r *http.Request) {

	// select last task
	row := db.QueryRow("SELECT * FROM tasks ORDER BY id DESC")

	switch err := row.Scan(&id, &title, &completed); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		task := models.Task{
			Id:         id,
			Title:      title,
			IsComplete: completed,
		}

		fmt.Println(task)
	default:
		log.Fatal(err)
	}
}

// push task to stack
func Add(w http.ResponseWriter, r *http.Request) {

	// getting request body
	title = r.FormValue("title")

	if _, err := db.Exec("INSERT INTO tasks (title, is_completed) VALUES ($1, false)", title); err != nil {
		// TODO: http.Redirect(w, r, "/", 409)
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", 301)
}

// pop task from stack
func Complete(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM tasks ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", 301)
}
