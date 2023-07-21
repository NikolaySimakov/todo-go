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

// show tasks
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

// show current (last) task
func Current(w http.ResponseWriter, r *http.Request) {
	row := db.QueryRow("SELECT * FROM tasks", 1)
	switch err := row.Scan(&id, &title, &completed); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id)
	default:
		log.Fatal(err)
	}

	task := models.Task{
		Id:         id,
		Title:      title,
		IsComplete: completed,
	}

	fmt.Println(task)
}

// push task to stack
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

// pop task from stack
func Complete(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM tasks ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
}

// cleat stack
func Delete(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
}
